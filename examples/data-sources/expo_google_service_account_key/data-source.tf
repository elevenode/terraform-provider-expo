# Look up a key already uploaded to EAS by the `project_id` of its JSON key file.
data "expo_google_service_account_key" "this" {
  project_identifier = "my-gcp-project"

  # Optional: pick one key when several belong to the same GCP project.
  client_email = "eas-submit@my-gcp-project.iam.gserviceaccount.com"
}
