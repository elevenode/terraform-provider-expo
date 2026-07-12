resource "expo_app" "expo_app" {
  name = "My App Name"
  slug = "my-app-slug"
}

data "expo_google_service_account_key" "this" {
  project_identifier = "..."
}

resource "expo_android_app_credentials" "this" {
  app_id                        = expo_app.this.id
  identifier                    = "com.example.myapp"
  google_service_account_key_id = data.expo_google_service_account_key.this.id
  fcm_key                       = "..."
  build_credentials {
    name        = "Default"
    keystore_id = "..."
  }
}
