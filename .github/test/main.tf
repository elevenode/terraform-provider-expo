terraform {
  required_providers {
    expo = {
      source = "elevenode/expo"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.6"
    }
  }
}

provider "expo" {}

# Branch/channel names are unique per app and are not reclaimed by a failed
# destroy, so a fixed name would wedge every later run.
resource "random_string" "update_name" {
  length  = 10
  special = false
  upper   = false
}

resource "expo_app" "this" {
  name = "Terraform Provider EAS"
  slug = "terraform-provider-expo"
}

resource "expo_app_variable" "this" {
  app_id       = expo_app.this.id
  name         = "API_KEY"
  value        = "my-api-key"
  visibility   = "PUBLIC"
  environments = ["DEVELOPMENT"]
}

resource "expo_account_variable" "this" {
  name         = "ACCOUNT_API_KEY"
  value        = "my-account-api-key"
  visibility   = "PUBLIC"
  environments = ["DEVELOPMENT"]
}

resource "expo_update_branch" "this" {
  app_id = expo_app.this.id
  name   = random_string.update_name.result
}

resource "expo_update_channel" "this" {
  app_id = expo_app.this.id
  name   = random_string.update_name.result

  branch_mapping = jsonencode({
    version = 0
    data = [{
      branchId           = expo_update_branch.this.id
      branchMappingLogic = "true"
    }]
  })
}

resource "expo_ios_app_identifier" "this" {
  identifier = local.bundle_identifier
}

resource "expo_ios_app_provisioning_profile" "this" {
  app_identifier_id = expo_ios_app_identifier.this.id
  base64            = var.PROVISIONING_PROFILE_BASE64
}

resource "expo_ios_app_credentials" "this" {
  app_id               = expo_app.this.id
  app_identifier_id    = expo_ios_app_identifier.this.id
  app_store_api_key_id = data.expo_app_store_api_key.this.id
  push_key_id          = data.expo_ios_push_key.this.id
  app_store {
    provisioning_profile_id = expo_ios_app_provisioning_profile.this.id
    certificate_id          = data.expo_ios_certificate.this.id
  }
}

resource "expo_android_app_credentials" "this" {
  app_id                        = expo_app.this.id
  identifier                    = local.bundle_identifier
  google_service_account_key_id = data.expo_google_service_account_key.this.id
  fcm_key                       = var.FCM_KEY
  build_credentials {
    name        = "Default"
    keystore_id = "67484c57-542f-48fc-a470-fa6703a3a6f5"
  }
}
