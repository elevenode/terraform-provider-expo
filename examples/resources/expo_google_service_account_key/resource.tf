resource "expo_google_service_account_key" "this" {
  key_json = file("${path.module}/play-service-account.json")
}
