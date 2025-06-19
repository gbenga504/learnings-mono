https://cloud.google.com/compute/docs/gcloud-compute#set_default_zone_and_region_in_your_local_client
https://sunitc.dev/2020/12/12/how-to-create-first-kubernetes-cluster-on-google-cloud-and-connect-to-it-from-local-machine-using-kubectl/

- gcloud config set project <project name> : To set the default project name
- gcloud config get-value compute/region: To get the default compute region
- gcloud config get-value compute/zone: To get the default compute zone
- gcloud config set compute/region <region name>: To set the default compute region
- gcloud config set compute/zone <zone name>: To set the default compute zone
- gcloud config unset compute/zone / gcloud config unset compute/region: To unset region or zone
- gcloud compute regions list / gcloud compute zones list: To list the regions or zones available on Google cloud
- gcloud config list: To list the config
- gcloud container clusters list: To list the clusters
- gcloud compute instances list: To list the available nodes
- kubectl config view: View the kubectl config
- gcloud container clusters get-credentials <cluster-name>: Config kubectl to connect to your google cloud cluster
- gcloud container clusters resize <cluster-name> --num-nodes=<size>: To scale down/up your nodes
