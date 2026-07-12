# Account-level environment variable. It belongs to the account configured on
# the provider (account_name), so no app/account id is needed.
resource "expo_account_variable" "this" {
  name         = "API_KEY"
  value        = "my-api-key"
  visibility   = "SECRET"
  environments = ["production"]
}
