resource "expo_ios_app_identifier" "this" {
  identifier = "my.app.identifier"
}

resource "expo_ios_app_provisioning_profile" "this" {
  app_identifier_id = expo_ios_app_identifier.this.id
  base64            = "..."
}
