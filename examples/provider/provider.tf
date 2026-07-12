terraform {
  required_providers {
    expo = {
      source = "elevenode/expo"
    }
  }
}

provider "expo" {
  token        = "..."
  account_name = "..."
}
