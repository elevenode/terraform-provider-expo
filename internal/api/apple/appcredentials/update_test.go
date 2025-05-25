package appcredentials

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/api/apple/appbuildcredentials"
	"github.com/fintreal/eas-sdk-go/internal/utils"
)

func TestUpdate(t *testing.T) {
	pushKeyId := "test-push-key-id"
	appStoreApiKeyId := "test-app-store-api-key-id"
	expectedData := &Data{
		Id:               "test-id",
		AppIdentifierId:  "test-app-identifier-id",
		AppId:            "test-app-id",
		PushKeyId:        &pushKeyId,
		AppStoreApiKeyId: &appStoreApiKeyId,
		BuildCredentials: []appbuildcredentials.Data{},
	}

	expectedVariables := map[string]any{
		"id":               expectedData.Id,
		"appStoreApiKeyId": expectedData.AppStoreApiKeyId,
		"pushKeyId":        expectedData.PushKeyId,
	}

	mockResponse := updateResponse{
		IosAppCredentials: updateIosAppCredentials{
			Data: data{
				Id: expectedData.Id,
				App: objWithId{
					Id: expectedData.AppId,
				},
				AppIdentifier: objWithId{
					Id: expectedData.AppIdentifierId,
				},
				PushKey: &objWithId{
					Id: pushKeyId,
				},
				AppStoreApiKey: &objWithId{
					Id: appStoreApiKeyId,
				},
				BuildCredentials: []buildCredentials{},
			},
		},
	}

	input := &UpdateData{
		Id:               expectedData.Id,
		AppStoreApiKeyId: &appStoreApiKeyId,
		PushKeyId:        &pushKeyId,
	}

	config := utils.TestConfig[UpdateData, Data, updateResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Update",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      updateQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
