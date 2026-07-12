data "expo_app_store_api_key" "this" {
  identifier = "..."
}

data "expo_ios_certificate" "this" {
  serial_number = "..."
}

resource "expo_app" "this" {
  name = "My App Name"
  slug = "my-app-slug"
}

resource "expo_ios_app_identifier" "this" {
  identifier = "my.app.identifier"
}

resource "expo_provisioning_profile" "this" {
  app_identifier_id = expo_ios_app_identifier.this.id
  base64            = "..."
}

resource "expo_ios_app_credentials" "this" {
  app_id               = expo_app.expo_app.id
  app_identifier_id    = expo_ios_app_identifier.this.id
  app_store_api_key_id = data.expo_app_store_api_key.this.id
  app_store {
    provisioning_profile_id = expo_provisioning_profile.this.id
    certificate_id          = data.expo_ios_certificate.this.id
  }
}
