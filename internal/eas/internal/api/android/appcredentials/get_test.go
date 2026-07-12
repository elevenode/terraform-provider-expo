package appcredentials

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/android/appbuildcredentials"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestGetByIdentifier(t *testing.T) {
	input := &GetData{
		Id:    "test-id",
		AppId: "test-app-id",
	}
	fcmKey := &fcmKey{
		Id:      "test-fcm-key-id",
		KeyJson: "test-fcm-key-json",
	}

	buildCredential := appbuildcredentials.Data{
		Id:               "test-build-credentials-id",
		Name:             "test-build-credentials-name",
		AppCredentialsId: input.Id,
		KeystoreId:       "test-keystore-id",
	}

	expectedData := &Data{
		Id:                        input.Id,
		AppId:                     input.AppId,
		Identifier:                "test-identifier",
		GoogleServiceAccountKeyId: "test-google-service-account-key-id",
		FCMKey:                    &fcmKey.KeyJson,
		FCMKeyId:                  &fcmKey.Id,
		BuildCredentials:          []appbuildcredentials.Data{buildCredential},
	}

	expectedVariables := map[string]any{
		"appId": input.AppId,
	}
	mockResponse := getResponse{
		AppByAppId: appByAppId{
			Data: []data{{
				Id:         expectedData.Id,
				Identifier: expectedData.Identifier,
				App: objWithId{
					Id: expectedData.AppId,
				},
				GoogleServiceAccountKey: objWithId{
					Id: expectedData.GoogleServiceAccountKeyId,
				},
				FCMKey: fcmKey,
				BuildCredentials: []buildCredentialsData{{
					Id:   buildCredential.Id,
					Name: buildCredential.Name,
					Keystore: objWithId{
						Id: buildCredential.KeystoreId,
					},
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
