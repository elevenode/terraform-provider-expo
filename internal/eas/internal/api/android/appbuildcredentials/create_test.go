package appbuildcredentials

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestCreate(t *testing.T) {
	expectedData := &Data{
		Id:               "test-id",
		Name:             "test-name",
		AppCredentialsId: "test-app-credentials-id",
		KeystoreId:       "test-keystore-id",
	}

	expectedVariables := map[string]any{
		"appCredentialsId": expectedData.AppCredentialsId,
		"keystoreId":       expectedData.KeystoreId,
		"name":             expectedData.Name,
	}

	mockResponse := createResponse{
		AndroidAppBuildCredentials: createAndroidAppBuildCredentials{
			Data: data{
				Id:   expectedData.Id,
				Name: expectedData.Name,
				Keystore: objWithId{
					Id: expectedData.KeystoreId,
				},
			},
		},
	}

	input := &CreateData{
		KeystoreId:       expectedData.KeystoreId,
		Name:             expectedData.Name,
		AppCredentialsId: expectedData.AppCredentialsId,
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
