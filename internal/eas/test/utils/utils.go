package utils

import (
	"math/rand"
	"os"

	"github.com/elevenode/terraform-provider-expo/internal/eas"
)

var Token = os.Getenv("EXPO_TOKEN")
var Client = eas.NewEASClient(Token)
var ImmutableProvisioningProfileBase64 = os.Getenv("IMMUTABLE_PROVISIONING_PROFILE_BASE64")
var FCMKey = os.Getenv("FCM_KEY")
var KeystoreBase64 = os.Getenv("KEYSTORE_BASE64")

// TEST IDs
//
// All live in account "expo-eas-sdk-go". Note the Immutable/Mutable prefixes describe
// whether a test mutates the fixture, NOT which app owns it: every ID below hangs off
// ImmutableAppId (test-app). MutableAppId is the only one naming a different app.

// Account "expo-eas-sdk-go".
var AccountId = getEnvOrDefault("ACCOUNT_ID", "c705db5b-7312-4d9e-baeb-65bb640b0888")

// Robot "integration-test" — the actor the CI EXPO_TOKEN authenticates as.
var MeId = "e013cf58-3bf8-4522-a4da-cde29b9fbc95"

// @expo-eas-sdk-go/test-app ("Integration Test Fixture"), read by the tests in this package.
// Owns every credential/variable fixture below. Not Terraform-managed.
var ImmutableAppId = "953ed82f-4ac7-47be-ab46-d9c7a1169fe6"

// @expo-eas-sdk-go/test-app-update. TestAppUpdate renames it, so only the slug is stable.
var MutableAppId = "976aa3ac-c3c7-47f5-a94e-2f3f0f75409d"

// On test-app: "TEST_ENVIRONMENT_VARIABLE" (preview).
var ImmutableAppVariableId = "955ef61a-b78c-43cc-bfa7-6da1a51e76e3"

// On test-app (production). TestAppVariableUpdate renames it, so the name is random.
var MutableAppVariableId = "2d90f56c-aead-41d8-8a76-dcbcb92415ac"

var ImmutableAccountVariableId = os.Getenv("IMMUTABLE_ACCOUNT_VARIABLE_ID")
var MutableAccountVariableId = os.Getenv("MUTABLE_ACCOUNT_VARIABLE_ID")

// Apple team TEST_IDENTIFIER / "TEST_NAME".
var ImmutableAppleTeamId = "cf4830de-3c1b-4594-ad73-20ca9ac901b7"

// Apple team TEST_UPDATE_TEAM. The update test renames it, so the team name is random.
var MutableAppleTeamId = "cc513f5d-366d-42d6-aa40-e8132c3c78a3"

// iOS credentials on test-app for bundle "test.app.identifier" (no apple team attached).
var ImmutableAppleAppCredentialsId = "11f2b3f8-ddad-4626-8984-2b96efb28d3c"

// iOS credentials on test-app for bundle "immutable.app.identifier" — despite the Mutable name.
var MutableAppCredentialsId = "c51879e7-cc4c-4f9c-b22d-41b571d6f0f0"

// Apple distribution certificate; its serial is the TF_VAR_IOS_CERTIFICATE_SERIAL_NUMBER secret.
var ImmutableCertificateId = "702635c5-3aa1-477c-83b6-bb66a1644aad"

// Apple provisioning profile for bundle "com.fintreal.testapp".
var ImmutableProvisioningProfileId = "8690db1b-c475-43d0-aa3f-67e103c96426"

// Android credentials on test-app for applicationIdentifier "my.test.app.identifier".
var ImmutableAndroidAppCredentialsId = "9994bb7d-e1d5-4d50-8429-addfa2f24f72"

// The "Default" build credentials under ImmutableAndroidAppCredentialsId.
var ImmutableAndroidAppBuildCredentialsId = "55c4357a-f0bf-43c3-aeff-9828a0b584d6"
var ImmutableAppIdentifierId = "ce3f1747-2362-45c8-b55d-c392e0b6e94b"
var ImmutableAppIdentifierName = "immutable.app.identifier"

// App Store Connect API key; its keyIdentifier is the TF_VAR_APP_STORE_API_KEY_IDENTIFIER secret.
var ImmutableAppStoreApiKeyId = "564e9d75-ff77-4860-92ee-7c0ab2066c82"

var MutableAppIdentifierId = "63b653a7-7472-4804-9da4-39b474c52167"
var MutableAppIdentifierName = "mutable.app.identifier"

// Keystore (keyAlias "android") under ImmutableAndroidAppBuildCredentialsId.
var ImmutableKeystoreId = "67484c57-542f-48fc-a470-fa6703a3a6f5"

const charset = "abcdefghijklmnopqrstuvwxyz"

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
