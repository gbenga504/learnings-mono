

provider "google" {
  project = local.project
  region  = "us-central1"
  zone    = "us-central1-c"
}

resource "google_container_cluster" "primary_cluster" {
  project                   = "total-triumph-428918-t6"
  name                      = "practice-production"
  location                  = "us-central1-c"
  network                   = "projects/total-triumph-428918-t6/global/networks/default"
  subnetwork                = "projects/total-triumph-428918-t6/regions/us-central1/subnetworks/default"
  default_max_pods_per_node = 110
  deletion_protection = false

  min_master_version = null

  ip_allocation_policy {
    cluster_ipv4_cidr_block  = "10.16.0.0/14"
    services_ipv4_cidr_block = "34.118.224.0/20"
  }
}

resource "google_container_node_pool" "node_pool" {
  name     = "practice-node-pool-1"
  location = var.zone
  cluster  = google_container_cluster.primary_cluster.name

  autoscaling {
    location_policy      = "BALANCED"
    max_node_count       = 4
    min_node_count       = 2
    total_max_node_count = 0

    total_min_node_count = 0
  }

  queued_provisioning {
    enabled = false
  }
}

# data "google_container_cluster" "primary_cluster" {
#   name = "practice-production"
#   location = "us-central1-c"
# }

# output "endpoint" {
#   value = data.google_container_cluster.primary_cluster.endpoint
# }

