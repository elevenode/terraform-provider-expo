# terraform-provider-eas

Terraform provider for managing Expo Application Services (EAS) apps, credentials, and environment variables via the Expo GraphQL API.

- Repo: `git@github.com:fintreal/terraform-provider-eas.git`
- Default branch: `main`
- Registry namespace: `fintreal/eas`

## Tech Stack

- Go 1.25
- Terraform Plugin SDK v2
- Expo GraphQL client, vendored in-repo at `internal/eas/` (formerly the standalone `fintreal/eas-sdk-go`, now folded in)
- GoReleaser for builds/releases
- GPG-signed releases to Terraform Registry

## Project Structure

```
main.go                          # Provider entrypoint
provider/provider.go             # Provider definition (schema, resources, data sources)
provider/<resource>/schema.go    # Resource/data source schema
provider/<resource>/operations/  # CRUD operations (create.go, read.go, update.go, delete.go)
internal/client/eas.go           # EAS API client wrapper
internal/eas/                    # Vendored Expo GraphQL client (public pkg `eas`)
internal/eas/internal/           # Client internals (api, graphql, utils) -- private
internal/eas/test/               # Client live integration tests (need EXPO_TOKEN)
examples/                        # Terraform example configs (used by tfplugindocs)
docs/                            # Auto-generated registry docs (do not edit manually)
.github/test/                    # Integration test Terraform configs
```

### Resources

- `eas_app` - Expo app
- `eas_app_variable` - App environment variable
- `eas_android_app_credentials` - Android app credentials
- `eas_ios_app_credentials` - iOS app credentials
- `eas_ios_app_identifier` - iOS app identifier (Bundle ID)
- `eas_ios_app_provisioning_profile` - iOS provisioning profile

### Data Sources

- `eas_app_store_api_key`, `eas_ios_certificate`, `eas_ios_push_key`, `eas_google_service_account_key`

## Key Commands

```bash
go build -o terraform-provider-eas   # Build the provider
go install .                         # Install locally for dev testing
go mod tidy                          # Tidy dependencies
```

Set `GOPRIVATE=github.com/fintreal/*` for private module access.

## CI/CD

- **PR / nightly** (test.yml): unit tests (`go test`, offline) + folded SDK integration tests (`internal/eas/test/`, live) + `terraform apply`/`destroy` against `.github/test/` configs
- **Push to main** (release.yml): Verify gate (`gofmt`/`go build`/`go vet`/unit tests) → auto-generates docs via pinned `tfplugindocs` → creates semver release → builds with GoReleaser
- Docs in `docs/` are auto-generated -- do not edit manually
- The `Integration Test` job needs these repo secrets (migrated from the old SDK repo): `EXPO_TOKEN`, `IMMUTABLE_PROVISIONING_PROFILE_BASE64`, `FCM_KEY`, `KEYSTORE_BASE64`

## Auth

Provider requires `EXPO_TOKEN` and `EXPO_ACCOUNT_NAME` environment variables (or inline config).

## Conventions

- Each resource/data source lives in its own package under `provider/`
- iOS resources are nested under `provider/ios/`, Android under `provider/android/`
- CRUD operations are split into separate files in an `operations/` sub-package
- Schema definition is in `schema.go` at the resource package root
- Resources return `*schema.Resource` via a `Resource()` function; data sources via `DataSource()`
- Commit messages use conventional commits style (e.g., `feat:`, `fix:`, `chore:`)
