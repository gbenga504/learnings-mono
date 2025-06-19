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
	"fmt"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	introv1alpha "github.com/busybox/api/v1alpha"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BusyboxReconciler reconciles a Busybox object
type BusyboxReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// Definitions to manage status conditions
const (
	// typeReadyBusybox represents the status of the Deployment reconciliation
	typeReadyBusybox = "Ready"
	// typeDegradedBusybox represents the status used when the custom resource is deleted and the finalizer operations are yet to occur.
	typeDegradedBusybox = "Degraded"
)

// +kubebuilder:rbac:groups=intro.github.com,resources=busyboxes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=intro.github.com,resources=busyboxes/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=intro.github.com,resources=busyboxes/finalizers,verbs=update
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Busybox object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.1/pkg/reconcile
func (r *BusyboxReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Fetch the Busybox instance
	// We want to check if the custom resource for the Busybox Kind has been applied to the cluster
	// If it has not been applied, then we return nil and stop the reconcilation
	busybox := &introv1alpha.Busybox{}
	err := r.Get(ctx, req.NamespacedName, busybox)

	if err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("busybox resource not found. Ignoring since object must have been deleted")
			return ctrl.Result{}, nil
		}

		// Maybe there was an error reading the object - In this case, we want to log an error
		// and retrigger reconciliation
		log.Error(err, "Failed to get busybox resource")
		return ctrl.Result{}, err
	}

	// We want to set the status to Unknown when no status is available
	if len(busybox.Status.Conditions) == 0 {
		meta.SetStatusCondition(&busybox.Status.Conditions, metav1.Condition{Type: typeReadyBusybox, Status: metav1.ConditionUnknown, Reason: "Reconciling", Message: "Starting reconciliation"})

		if err = r.Status().Update(ctx, busybox); err != nil {
			log.Error(err, "Fauled to update busybox status")
			return ctrl.Result{}, err
		}

		// After updating the status, we want to refetch the resource so it can be used later in this function.
		// If we don't refetch the resource, we risk triggering an error that says "the object has been modified, please apply
		// your changes to the latest version and try again" and triggering a requeue cycle for the reconciliation
		if err = r.Get(ctx, req.NamespacedName, busybox); err != nil {
			log.Error(err, "Could not re-fetch the busybox resource")
			return ctrl.Result{}, err
		}
	}

	// Check if the deployment already exists, if not create a new one
	found := &appsv1.Deployment{}
	// For the key object, we use a new namespacedName object because we want to find a deployment
	// in the same namespace as our busybox custom resource and most especially, a deployment whose name
	// is the same with our buysbox custom resource. It is because of this reason, we don't rely on req.NamespacedName
	// since the Name property could be different from our busybox resource name
	err = r.Get(ctx, types.NamespacedName{Name: busybox.Name, Namespace: busybox.Namespace}, found)

	if err != nil && apierrors.IsNotFound(err) {
		// Define a new deployment
		deployment, err := r.deploymentForBusybox(busybox)

		if err != nil {
			log.Error(err, "Failed to define new deployment resource for busybox")

			// The following implementation will update the status
			meta.SetStatusCondition(&busybox.Status.Conditions,
				metav1.Condition{Type: typeReadyBusybox,
					Status: metav1.ConditionFalse, Reason: "Reconciling",
					Message: fmt.Sprintf("Failed to create Deployment for the custom resources (%s): (%s)", busybox.Name, err),
				})

			if err := r.Status().Update(ctx, busybox); err != nil {
				log.Error(err, "Failed to update busybox status")
				return ctrl.Result{}, err
			}

			return ctrl.Result{}, err
		}

		log.Info("Creating a new deployment", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)

		if err = r.Create(ctx, deployment); err != nil {
			log.Error(err, "Failed to create new deployment", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)

			return ctrl.Result{}, err
		}

		// Deployment was created successfully
		// Here we need to requeue the reconciliation so that we can ensure that the new state matches the desired state by the Custom resource
		// and move forward with the next operations
		return ctrl.Result{RequeueAfter: time.Minute}, nil
	} else if err != nil {
		log.Error(err, "Failed to get deployment")

		// Lets return the err so that we can retrigger reconciliation
		return ctrl.Result{}, err
	}

	// The CRD API defines that Busybox type specify a size field in its spec
	// Hence during reconciliation, we have to check that the deployment replicas matches the Size specified in the Busybox custom resource
	size := busybox.Spec.Size
	if *found.Spec.Replicas != size {
		found.Spec.Replicas = &size

		if err = r.Update(ctx, found); err != nil {
			log.Error(err, "Failed to update Deployment",
				"Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)

			// Once again we need to refetch the busybo resources before updating the status
			if err = r.Get(ctx, req.NamespacedName, busybox); err != nil {
				log.Error(err, "Failed to re-fetch busybox")
				return ctrl.Result{}, err
			}

			// The following will update the status
			meta.SetStatusCondition(&busybox.Status.Conditions, metav1.Condition{Type: typeReadyBusybox,
				Status: metav1.ConditionFalse, Reason: "Resizing",
				Message: fmt.Sprintf("Failed to update the size for the custom resource (%s): (%s)", busybox.Name, err)})

			if err = r.Update(ctx, busybox); err != nil {
				log.Error(err, "Failed to update Busybox status")
				return ctrl.Result{}, err
			}

			return ctrl.Result{}, err
		}

		// Now that we have updated the replicas, we need to requeue the reconciliation
		// to make sure that everything is up to date before continuing other operation
		return ctrl.Result{Requeue: true}, nil
	}

	// Here the replica of our busybox deployment matches with the size specified the Busybox CR spec
	// Hence we want to update the status of the Busybox CR
	meta.SetStatusCondition(&busybox.Status.Conditions, metav1.Condition{Type: typeReadyBusybox, Status: metav1.ConditionTrue,
		Reason: "Reconciling", Message: fmt.Sprintf("Deployment for custom resource (%s) with %d replicas created successfully", busybox.Name, size)})

	if err = r.Update(ctx, busybox); err != nil {
		log.Error(err, "Failed to update the status of Busybox CR")

		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *BusyboxReconciler) deploymentForBusybox(busybox *introv1alpha.Busybox) (*appsv1.Deployment, error) {
	replicas := busybox.Spec.Size
	image := "busybox:latest"

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      busybox.Name,
			Namespace: busybox.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app.kubernetes.io/name": "project"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app.kubernetes.io/name": "project"},
				},
				Spec: corev1.PodSpec{
					SecurityContext: &corev1.PodSecurityContext{
						RunAsNonRoot: &[]bool{true}[0],
						SeccompProfile: &corev1.SeccompProfile{
							Type: corev1.SeccompProfileTypeRuntimeDefault,
						},
					},
					Containers: []corev1.Container{{
						Image:           image,
						Name:            "busybox",
						ImagePullPolicy: corev1.PullIfNotPresent,
						// Ensure restrictive context for the container
						// More info: https://kubernetes.io/docs/concepts/security/pod-security-standards/#restricted
						SecurityContext: &corev1.SecurityContext{
							RunAsNonRoot:             &[]bool{true}[0],
							RunAsUser:                &[]int64{1001}[0],
							AllowPrivilegeEscalation: &[]bool{false}[0],
							Capabilities: &corev1.Capabilities{
								Drop: []corev1.Capability{
									"ALL",
								},
							},
						},

						Ports: []corev1.ContainerPort{{
							ContainerPort: 8000,
							Name:          "busybox",
						}},

						Command: []string{"sh", "-c", "echo 'Hello world' && sleep 3600"},
					}},
				},
			},
		},
	}

	// Set the ownerRef for the Deployment
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/owners-dependents/
	if err := ctrl.SetControllerReference(busybox, deployment, r.Scheme); err != nil {
		return nil, err
	}

	return deployment, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *BusyboxReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&introv1alpha.Busybox{}).
		Owns(&appsv1.Deployment{}).
		Named("busybox").
		Complete(r)
}
