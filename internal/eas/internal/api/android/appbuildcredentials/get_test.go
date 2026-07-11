package appbuildcredentials

import (
	"testing"

	"terraform-provider-eas/internal/eas/internal/utils"
)

func TestGet(t *testing.T) {
	input := &GetData{
		Id:               "test-id",
		AppId:            "test-app-id",
		AppCredentialsId: "test-app-credentials-id",
	}
	expectedData := &Data{
		Id:               input.Id,
		KeystoreId:       "test-keystore-id",
		AppCredentialsId: input.AppCredentialsId,
		Name:             "test-name",
	}

	expectedVariables := map[string]any{"appId": input.AppId}

	mockResponse := getResponse{
		AppByAppId: appByAppId{
			AndroidAppCredentials: []androidAppCredentials{{
				Id: expectedData.AppCredentialsId,
				Data: []data{{
					Id:       expectedData.Id,
					Keystore: objWithId{Id: expectedData.KeystoreId},
					Name:     expectedData.Name,
				}},
			}},
		},
	}

	config := utils.TestConfig[GetData, Data, getResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Get",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      getQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
