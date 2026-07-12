package appcredentials

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/android/appbuildcredentials"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestCreate(t *testing.T) {
	expectedData := &Data{
		Id:                        "test-id",
		AppId:                     "test-app-id",
		Identifier:                "test-identifier",
		GoogleServiceAccountKeyId: "test-google-service-account-key-id",
		BuildCredentials:          []appbuildcredentials.Data{},
	}

	expectedVariables := map[string]any{
		"appId":                     expectedData.AppId,
		"identifier":                expectedData.Identifier,
		"googleServiceAccountKeyId": expectedData.GoogleServiceAccountKeyId,
	}

	mockResponse := createResponse{
		AndroidAppCredentials: createAndroidAppCredentials{
			Data: data{
				Id: expectedData.Id,
				App: objWithId{
					Id: expectedData.AppId,
				},
				Identifier: expectedData.Identifier,
				GoogleServiceAccountKey: objWithId{
					Id: expectedData.GoogleServiceAccountKeyId,
				},
			},
		},
	}

	input := &CreateData{
		AppId:                     expectedData.AppId,
		Identifier:                expectedData.Identifier,
		GoogleServiceAccountKeyId: expectedData.GoogleServiceAccountKeyId,
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
