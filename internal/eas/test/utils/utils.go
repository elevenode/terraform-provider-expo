package utils

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/elevenode/terraform-provider-expo/internal/eas"
)

var Token = os.Getenv("EXPO_TOKEN")
var Client = eas.NewEASClient(Token)
var ImmutableProvisioningProfileBase64 = os.Getenv("IMMUTABLE_PROVISIONING_PROFILE_BASE64")
var FCMKey = os.Getenv("FCM_KEY")
var KeystoreBase64 = os.Getenv("KEYSTORE_BASE64")

// TEST IDs

var AccountId = getEnvOrDefault("ACCOUNT_ID", "c705db5b-7312-4d9e-baeb-65bb640b0888")
var MeId = "e013cf58-3bf8-4522-a4da-cde29b9fbc95"

var ImmutableAppId = "953ed82f-4ac7-47be-ab46-d9c7a1169fe6"
var MutableAppId = "976aa3ac-c3c7-47f5-a94e-2f3f0f75409d"

var ImmutableAppVariableId = "955ef61a-b78c-43cc-bfa7-6da1a51e76e3"
var MutableAppVariableId = "2d90f56c-aead-41d8-8a76-dcbcb92415ac"

var ImmutableAccountVariableId = os.Getenv("IMMUTABLE_ACCOUNT_VARIABLE_ID")
var MutableAccountVariableId = os.Getenv("MUTABLE_ACCOUNT_VARIABLE_ID")

var ImmutableAppleTeamId = "cf4830de-3c1b-4594-ad73-20ca9ac901b7"
var MutableAppleTeamId = "cc513f5d-366d-42d6-aa40-e8132c3c78a3"

var ImmutableAppleAppCredentialsId = "11f2b3f8-ddad-4626-8984-2b96efb28d3c"

var MutableAppCredentialsId = "c51879e7-cc4c-4f9c-b22d-41b571d6f0f0"

var ImmutableCertificateId = "702635c5-3aa1-477c-83b6-bb66a1644aad"

var ImmutableProvisioningProfileId = "8690db1b-c475-43d0-aa3f-67e103c96426"

var ImmutableAndroidAppCredentialsId = "9994bb7d-e1d5-4d50-8429-addfa2f24f72"
var ImmutableAndroidAppBuildCredentialsId = "55c4357a-f0bf-43c3-aeff-9828a0b584d6"
var ImmutableAppIdentifierId = "ce3f1747-2362-45c8-b55d-c392e0b6e94b"
var ImmutableAppIdentifierName = "immutable.app.identifier"

var ImmutableAppStoreApiKeyId = "564e9d75-ff77-4860-92ee-7c0ab2066c82"

var MutableAppIdentifierId = "63b653a7-7472-4804-9da4-39b474c52167"
var MutableAppIdentifierName = "mutable.app.identifier"

var ImmutableKeystoreId = "67484c57-542f-48fc-a470-fa6703a3a6f5"

const charset = "abcdefghijklmnopqrstuvwxyz"

const (
	AppDeletionTimeout        = 2 * time.Minute
	backgroundJobPollInterval = 2 * time.Second
)

// WaitForBackgroundJob polls a background job receipt until it settles or the timeout expires.
func WaitForBackgroundJob(id string, timeout time.Duration) (*eas.BackgroundJobData, error) {
	deadline := time.Now().Add(timeout)

	for {
		receipt, err := Client.BackgroundJob.Get(id)
		if err != nil {
			return nil, err
		}
		if receipt == nil {
			return nil, fmt.Errorf("background job %s not found", id)
		}

		switch {
		case receipt.State == eas.BackgroundJobStateSuccess:
			return receipt, nil
		case receipt.State == eas.BackgroundJobStateFailure && !receipt.WillRetry:
			return receipt, fmt.Errorf("background job %s failed: %s %s", id, receipt.ErrorCode, receipt.ErrorMessage)
		}

		if time.Now().After(deadline) {
			return receipt, fmt.Errorf("timed out waiting for background job %s, last state %s", id, receipt.State)
		}

		time.Sleep(backgroundJobPollInterval)
	}
}

// DeleteApp deletes an app and waits for the deletion to finish, so tests don't leak projects.
func DeleteApp(id string) error {
	receipt, err := Client.App.Delete(id)
	if err != nil {
		return err
	}
	if receipt == nil {
		return fmt.Errorf("no background job receipt returned for the deletion of app %s", id)
	}

	_, err = WaitForBackgroundJob(receipt.Id, AppDeletionTimeout)
	return err
}

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
