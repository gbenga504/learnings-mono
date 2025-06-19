locals {
  project = "total-triumph-428918-t6"
}

variable "zone" {
  type = string
  default = "us-central1-c"
  description = "This is the zone for our cluster"
  sensitive = false
}
