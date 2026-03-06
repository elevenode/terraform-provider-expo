# eas-sdk-go

Expo Application Services (EAS) SDK for Go -- manages Expo apps, credentials, and environment variables via the Expo GraphQL API.

- **Repo:** `git@github.com:fintreal/eas-sdk-go.git`
- **Default branch:** main

## Tech Stack

- Go 1.24.1
- GraphQL client: `machinebox/graphql`
- Testing: `stretchr/testify`
- CI/CD: GitHub Actions (test on PR, auto-release with semver on merge to main)

## Project Structure

```
eas/              # Public API surface (EASClient, exported type aliases)
internal/
  api/            # API domain services, each with CRUD operations + types
    account/      # Expo account
    app/          # Expo app management
    appvariable/  # Environment variables
    apple/        # Apple credentials (team, certs, profiles, push keys, app store keys)
    android/      # Android credentials (keystore, FCM key, Google service account key)
    me/           # Current user
  graphql/        # GraphQL client wrapper and mocks
  utils/          # Shared utilities
test/             # Integration tests (run against live Expo API with EXPO_TOKEN)
```

## Key Commands

```bash
# Unit tests
go test ./internal/... -count=1 -cover

# Integration tests (requires EXPO_TOKEN env var)
go test ./test/... -count=1 -cover
```

## Conventions

- Public package is `eas/`; all domain logic lives in `internal/` (unexported)
- Type aliases in `eas/types.go` re-export internal types for public consumption
- Each API domain follows a consistent pattern: `types.go` for structs, separate files per operation
- Unit tests use mocked GraphQL client; integration tests in `test/` hit the live API
- Releases are automated via semver tagging on merge to main
