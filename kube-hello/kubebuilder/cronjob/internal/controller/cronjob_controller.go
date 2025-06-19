/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"sort"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ref "k8s.io/client-go/tools/reference"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	batchv1 "github.com/cronjob/api/v1"
	kbatch "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type realClock struct{}

func (_ realClock) Now() time.Time { return time.Now() }

// Clock knows how to get the current time.
// It can be used to fake out timing for testing.
type Clock interface {
	Now() time.Time
}

// CronJobReconciler reconciles a CronJob object
type CronJobReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Clock
}

var jobOwnerKey = ".metadata.controller"
var scheduledTimeAnnotation = "batch.github.com/scheduled-at"

// find the active list of jons
var activeJobs []*kbatch.Job
var successfulJobs []*kbatch.Job
var failedJobs []*kbatch.Job
var lastScheduleTime *time.Time // find the last run so we can update the status

func isJobFinished(job *kbatch.Job) (bool, kbatch.JobConditionType) {
	for _, c := range job.Status.Conditions {
		if (c.Type == kbatch.JobComplete || c.Type == kbatch.JobFailed) && c.Status == corev1.ConditionTrue {
			return true, c.Type
		}
	}

	return false, ""
}

func getScheduledTimeForJob(job *kbatch.Job) (*time.Time, error) {
	timeRaw := job.Annotations[scheduledTimeAnnotation]

	if len(timeRaw) == 0 {
		return nil, nil
	}

	timeParsed, err := time.Parse(time.RFC3339, timeRaw)

	if err != nil {
		return nil, err
	}

	return &timeParsed, nil
}

// +kubebuilder:rbac:groups=batch.github.com,resources=cronjobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch.github.com,resources=cronjobs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=batch.github.com,resources=cronjobs/finalizers,verbs=update
// +kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch,resources=jobs/status,verbs=get

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CronJob object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.1/pkg/reconcile
func (r *CronJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	cronjob := &batchv1.CronJob{}

	if err := r.Get(ctx, req.NamespacedName, cronjob); err != nil {
		log.Error(err, "unable to fetch CronJob")

		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	childJobs := &kbatch.JobList{}
	// We need to list all the child Jobs that belongs to this CronJob
	// We get the child jobs using
	// 	1. client.InNamespace(req.Namespace) to get all the Jobs in the same namespace as the CronJob
	// 	2. client.MatchingFields{jobOwnerKey: req.Name} to get all the Jobs that have a controller owner reference to the CronJob
	//	   i.e we want all jobs that have a jobOwnerKey(.metadata.controller) field set to the name of the CronJob
	//	   This acts as a form of indexing for our jobs. Later in this document we will configure the manager to actually index this field
	if err := r.List(ctx, childJobs, client.InNamespace(req.Namespace), client.MatchingFields{jobOwnerKey: req.Name}); err != nil {
		log.Error(err, "unable to list child Jobs")

		return ctrl.Result{}, err
	}

	// We want to group the jobs into active, successful and failed jobs
	for _, job := range childJobs.Items {
		_, finishedType := isJobFinished(&job)

		switch finishedType {
		case "":
			activeJobs = append(activeJobs, &job)
		case kbatch.JobComplete:
			successfulJobs = append(successfulJobs, &job)
		case kbatch.JobFailed:
			failedJobs = append(failedJobs, &job)
		}

		// We'll store the launch time in an annotation, so we'll reconstitute that from
		// the active jobs themselves.
		scheduledTime, err := getScheduledTimeForJob(&job)

		if err != nil {
			log.Error(err, "unable to parse schedule time for child job", "job", &job)

			continue
		}

		if scheduledTime != nil {
			if lastScheduleTime == nil || lastScheduleTime.Before(*scheduledTime) {
				lastScheduleTime = scheduledTime
			}
		}
	}

	// Here we want to get the last schedule time for the cronjob i.e a job that would take the longest time before running
	if lastScheduleTime != nil {
		cronjob.Status.LastScheduleTime = &metav1.Time{Time: *lastScheduleTime}
	} else {
		cronjob.Status.LastScheduleTime = nil
	}

	// We want to update the cronJob status subresource CR with the active jobs
	// However, to do that we need to convert the active jobs to ObjectReferences
	cronjob.Status.Active = nil
	for _, activeJob := range activeJobs {
		jobRef, err := ref.GetReference(r.Scheme, activeJob)

		if err != nil {
			log.Error(err, "unable to make reference to active job", "job", activeJob)
			continue
		}

		cronjob.Status.Active = append(cronjob.Status.Active, *jobRef)
	}

	// We want to log at a debugging level. We also use key-value pairs so it is easy to filter and query log lines
	log.V(1).Info("job count", "active jobs", len(activeJobs), "successful jobs", len(successfulJobs), "failed jobs", len(failedJobs))

	// We want to update the status subresource with all the changes made to it so far
	// the r.Status.Update command ignores changes to the Spec field so we don't have conflict with other proesses (if any) trying to update the Spec field
	if err := r.Status().Update(ctx, cronjob); err != nil {
		log.Error(err, "unable to update CronJob status")

		return ctrl.Result{}, err
	}

	// We want to clean tup old jobs according to the history limits
	// NB: deleting these are "best effort" -- if we fail on a particular one,
	// we won't requeue just to finish the deleting.

	// We first sort the failed jobs by their start time. From the oldest to the newest
	if cronjob.Spec.FailedJobsHistoryLimit != nil {
		sort.Slice(failedJobs, func(i int, j int) bool {
			if failedJobs[i].Status.StartTime == nil {
				return failedJobs[j].Status.StartTime != nil
			}
			return failedJobs[i].Status.StartTime.Before(failedJobs[j].Status.StartTime)
		})

		for i, job := range failedJobs {
			// then we delete the jobs that are older than the history limit
			if int32(i) >= int32(len(failedJobs))-*cronjob.Spec.FailedJobsHistoryLimit {
				break
			}

			// We delete the job and all its dependents in the background so we don't block the controller and we don't have orphaned resources
			if err := r.Delete(ctx, job, client.PropagationPolicy(metav1.DeletePropagationBackground)); err != nil {
				log.Error(err, "unable to delete old failed job", "job", job)
			} else {
				// We create a log at the root level
				log.V(0).Info("deleted old failed job", "job", job)
			}
		}
	}

	// We first sort the successful jobs by their start time. From the oldest to the newest
	if cronjob.Spec.SuccessfulJobsHistoryLimit != nil {
		sort.Slice(successfulJobs, func(i int, j int) bool {
			if successfulJobs[i].Status.StartTime == nil {
				return successfulJobs[j].Status.StartTime != nil
			}
			return successfulJobs[i].Status.StartTime.Before(successfulJobs[j].Status.StartTime)
		})

		for i, job := range successfulJobs {
			// then we delete the jobs that are older than the history limit
			if int32(i) >= int32(len(successfulJobs))-*cronjob.Spec.SuccessfulJobsHistoryLimit {
				break
			}

			// We delete the job and all its dependents in the background so we don't block the controller and we don't have orphaned resources
			if err := r.Delete(ctx, job, client.PropagationPolicy(metav1.DeletePropagationBackground)); err != nil {
				log.Error(err, "unable to delete old successful job", "job", job)
			} else {
				// We create a log at the root level
				log.V(0).Info("deleted old successful job", "job", job)
			}
		}
	}

	// if this object is suspended, we don’t want to run any jobs, so we’ll stop now.
	// This is useful if something’s broken with the job we’re running and we want to pause runs to investigate without deleting the object
	if cronjob.Spec.Suspend != nil && *cronjob.Spec.Suspend {
		log.V(1).Info("cronjob suspended, skipping")

		return ctrl.Result{}, nil
	}

	// At this point, we need to get the next scheduled run or a run we have not scheduled yet

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CronJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1.CronJob{}).
		Named("cronjob").
		Complete(r)
}
