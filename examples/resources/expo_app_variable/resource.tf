resource "expo_app" "this" {
  name = "My App Name"
  slug = "my-app-slug"
}

resource "expo_app_variable" "this" {
  app_id       = expo_app.this.id
  name         = "API_URL"
  value        = "http://example.com/api"
  visibility   = "PUBLIC"
  environments = ["DEVELOPMENT"]
}

# File variable, e.g. the Firebase config consumed by app.config.js as
# process.env.GOOGLE_SERVICES_JSON (EAS writes the file and sets the path).
resource "expo_app_variable" "google_services_json" {
  app_id       = expo_app.this.id
  name         = "GOOGLE_SERVICES_JSON"
  type         = "FILE_BASE64"
  file_name    = "google-services.json"
  value        = filebase64("${path.module}/google-services.json")
  visibility   = "SECRET"
  environments = ["preview", "production"]
}

# The value can also come from another resource that exposes base64 file
# content, such as google_firebase_android_app_config.config_file_contents:
#   value = data.google_firebase_android_app_config.this.config_file_contents
