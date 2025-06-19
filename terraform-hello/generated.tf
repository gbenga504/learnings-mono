# __generated__ by Terraform
# Please review these resources and move them into your main configuration files.

# __generated__ by Terraform from "AudiJs"
resource "github_repository" "audijs_repo" {
  allow_auto_merge                        = false
  allow_merge_commit                      = true
  allow_rebase_merge                      = true
  allow_squash_merge                      = true
  allow_update_branch                     = false
  archive_on_destroy                      = null
  archived                                = false
  auto_init                               = false
  delete_branch_on_merge                  = false
  description                             = "Text-To-Speech library for javascript based mainly on manipulating the browsers speech API and providing workarounds if not available"
  gitignore_template                      = null
  has_discussions                         = false
  has_downloads                           = true
  has_issues                              = true
  has_projects                            = true
  has_wiki                                = true
  homepage_url                            = null
  ignore_vulnerability_alerts_during_read = null
  is_template                             = false
  license_template                        = null
  merge_commit_message                    = "PR_TITLE"
  merge_commit_title                      = "MERGE_MESSAGE"
  name                                    = "AudiJs"
  squash_merge_commit_message             = "COMMIT_MESSAGES"
  squash_merge_commit_title               = "COMMIT_OR_PR_TITLE"
  topics                                  = []
  visibility                              = "public"
  vulnerability_alerts                    = true
  web_commit_signoff_required             = false
  security_and_analysis {
    secret_scanning {
      status = "disabled"
    }
    secret_scanning_push_protection {
      status = "disabled"
    }
  }
}
