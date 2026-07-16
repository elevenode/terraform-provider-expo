resource "expo_update_branch" "production" {
  app_id = expo_app.this.id
  name   = "production"
}
