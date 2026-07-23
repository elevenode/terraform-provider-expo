# Expo Application Services (EAS) Terraform Provider

[![Terraform Registry](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fregistry.terraform.io%2Fv1%2Fproviders%2Felevenode%2Fexpo&query=%24.version&label=terraform%20registry&color=7B42BC&logo=terraform)](https://registry.terraform.io/providers/elevenode/expo/latest)
[![License](https://img.shields.io/github/license/elevenode/terraform-provider-expo)](./LICENSE)

Manage [Expo Application Services (EAS)](https://expo.dev/eas) as code: EAS apps, credentials, environment variables, and update channels. Backed by the Expo EAS GraphQL API.

## Installation

Add the provider to your Terraform configuration:

```hcl
terraform {
  required_providers {
    expo = {
      source = "elevenode/expo"
    }
  }
}

provider "expo" {
  access_token = var.expo_access_token # or set EXPO_TOKEN
  account_name = "your-account-name"   # or set EXPO_ACCOUNT_NAME
}
```

Then run `terraform init` to download the provider from the [Terraform Registry](https://registry.terraform.io/providers/elevenode/expo/latest).

## Authentication

The provider requires an Expo access token and account name.

1. **Access Token**
   - Log into your Expo account at https://expo.dev
   - Go to your account settings
   - Create a new access token under the "Access Tokens" section
   - Provide it via the `access_token` argument or the `EXPO_TOKEN` environment variable

2. **Account Name**
   - Your Expo account username (not email)
   - Found in your account settings or in the URL when logged into expo.dev (e.g. `https://expo.dev/your-account-name`)
   - For organizations, use the organization's account name
   - Provide it via the `account_name` argument or the `EXPO_ACCOUNT_NAME` environment variable

## Documentation

Full resource and data-source documentation is on the [Terraform Registry](https://registry.terraform.io/providers/elevenode/expo/latest/docs).

Runnable examples live in the [examples directory](./examples):

- [Resource Examples](./examples/resources) — managing EAS apps, credentials, and variables
- [Data Source Examples](./examples/data-sources) — querying EAS apps and variables
- [Provider Configuration Examples](./examples/provider) — additional provider configuration

## Contributing

Issues and pull requests are welcome. Please open an issue to discuss significant changes before submitting a PR.

## Security

Please report vulnerabilities privately as described in [SECURITY.md](./SECURITY.md).

## License

Apache 2.0 — see [LICENSE](./LICENSE).
