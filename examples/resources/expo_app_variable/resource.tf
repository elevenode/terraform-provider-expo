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
