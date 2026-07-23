# Security Policy

## Supported Versions

The latest released version on the [Terraform Registry](https://registry.terraform.io/providers/elevenode/expo/latest) is supported. Please upgrade to the latest version before reporting an issue.

## Reporting a Vulnerability

Please do **not** open a public issue for security vulnerabilities.

Instead, report them privately using GitHub's [private vulnerability reporting](https://github.com/elevenode/terraform-provider-expo/security/advisories/new) ("Report a vulnerability" under the Security tab).

Please include:

- A description of the vulnerability and its impact
- Steps to reproduce
- Affected version(s)

We aim to acknowledge reports within a few business days and will keep you updated on remediation progress.

## Handling Credentials

This provider requires an Expo access token (`EXPO_TOKEN`) and account name. Treat these as secrets:

- Never commit tokens or credentials to source control.
- Prefer environment variables or a secrets manager over inline configuration.
- Store Terraform state securely — it may contain sensitive values.
