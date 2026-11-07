resource "expo_app" "this" {
  name = "My App Name"
  slug = "my-app-slug"
}

# Upload a Google service account key for Play Store submissions.
# To use a key that is already in EAS, use the data source instead:
#   data "expo_google_service_account_key" "this" { project_identifier = "my-gcp-project" }
resource "expo_google_service_account_key" "submissions" {
  key_json = file("${path.module}/play-service-account.json")
}

resource "expo_android_keystore" "release" {
  keystore_base64   = filebase64("${path.module}/release.jks")
  keystore_password = var.keystore_password
  key_alias         = "upload"
  key_password      = var.key_password
}

resource "expo_android_app_credentials" "this" {
  app_id                        = expo_app.this.id
  identifier                    = "com.example.myapp"
  google_service_account_key_id = expo_google_service_account_key.submissions.id

  # Optional: service account key JSON used for FCM V1 push notifications.
  fcm_key = file("${path.module}/fcm-service-account.json")

  build_credentials {
    name        = "Default"
    keystore_id = expo_android_keystore.release.id
  }
}
