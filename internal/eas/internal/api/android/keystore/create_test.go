package keystore

import (
	"testing"

	"terraform-provider-eas/internal/eas/internal/utils"
)

func TestCreate(t *testing.T) {
	input := &CreateData{
		AccountId:        "test-account-id",
		KeystoreBase64:   "test-keystore-base64",
		KeyAlias:         "test-key-alias",
		KeyPassword:      "test-key-password",
		KeystorePassword: "test-keystore-password",
	}

	expectedData := &Data{
		Id:               "test-id",
		KeyAlias:         input.KeyAlias,
		KeyPassword:      input.KeyPassword,
		Keystore:         "test-keystore",
		KeystorePassword: input.KeystorePassword,
	}

	expectedVariables := map[string]any{
		"accountId":        input.AccountId,
		"keystoreBase64":   input.KeystoreBase64,
		"keyAlias":         input.KeyAlias,
		"keyPassword":      input.KeyPassword,
		"keystorePassword": input.KeystorePassword,
	}

	mockResponse := createResponse{
		AndroidKeystore: createAndroidKeystore{
			Data: *expectedData,
		},
	}

	config := utils.TestConfig[CreateData, Data, createResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Create",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      createQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
