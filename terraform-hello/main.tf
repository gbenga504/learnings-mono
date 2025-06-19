variable "token" {
  type    = string
  default = ""
}

# Since we do not provide a default, the user will be prompted to enter a value in the CLI
# We can also supply the value using ====> terraform plan -var="github_repo_name=VALUE"
variable "github_repo_name" {
  type = string
  validation {
    condition     = length(var.github_repo_name) <= 10
    error_message = "The repo name cannot be more than 10 characters"
  }
}

# We hooked this up to tfvars file, hence when running, we need to supply a tfvars file 
# terraform plan -var-file="prod.tfvars"
variable "github_repo_description" {
  type = string
}

locals {
  owner = "gbenga504"
}

provider "github" {
  token = var.token
  owner = local.owner
}

resource "github_repository" "terraform_test_repo" {
  name        = var.github_repo_name
  description = "${var.github_repo_description}-DESCRIPTION"

  visibility = "public"
}

output "github_url" {
  value = github_repository.terraform_test_repo.html_url
}

output "github_repo_id" {
  value = github_repository.terraform_test_repo.repo_id
}

# module "google" {
#   source = "./google"
# }
