package fcmkey

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestCreate(t *testing.T) {

	input := &CreateData{
		AccountId:        "test-account-id",
		AppCredentialsId: "test-app-credentials-id",
		KeyJson:          "test-key-json",
	}

	expectedData := &Data{
		Id:      "test-id",
		KeyJson: input.KeyJson,
	}

	expectedVariables := map[string]any{
		"appCredentialsId": input.AppCredentialsId,
		"keyJson":          input.KeyJson,
		"accountId":        input.AccountId,
	}

	mockResponse := createResponse{
		AndroidAppCredentials: androidAppCredentials{
			GoogleServiceAccountKeyData: googleServiceAccountKeyForFcmV1{
				Data: Data{
					Id:      expectedData.Id,
					KeyJson: expectedData.KeyJson,
				},
			},
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
