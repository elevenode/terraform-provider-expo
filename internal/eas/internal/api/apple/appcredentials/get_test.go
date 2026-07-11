package appcredentials

import (
	"testing"

	"terraform-provider-eas/internal/eas/internal/api/apple/appbuildcredentials"
	"terraform-provider-eas/internal/eas/internal/utils"
)

func TestGetByIdentifier(t *testing.T) {
	pushKeyId := "test-push-key-id"
	appStoreApiKeyId := "test-app-store-api-key-id"
	input := &GetData{
		Id:    "test-id",
		AppId: "test-app-id",
	}
	buildCredential := appbuildcredentials.Data{
		Id:                    "test-build-credential-id",
		DistributionType:      "test-distribution-type",
		CertificateId:         "test-certificate-id",
		ProvisioningProfileId: "test-provisioning-profile-id",
		AppCredentialsId:      input.Id,
	}
	expectedData := &Data{
		Id:               input.Id,
		AppIdentifierId:  "test-app-identifier-id",
		AppId:            input.AppId,
		PushKeyId:        &pushKeyId,
		AppStoreApiKeyId: &appStoreApiKeyId,
		BuildCredentials: []appbuildcredentials.Data{buildCredential},
	}

	expectedVariables := map[string]any{
		"id":    input.Id,
		"appId": input.AppId,
	}
	mockResponse := getResponse{
		AppByAppId: appByAppId{
			Data: []data{{
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
				BuildCredentials: []buildCredentials{{
					Id:               buildCredential.Id,
					DistributionType: buildCredential.DistributionType,
					AppCredentials: objWithId{
						Id: buildCredential.AppCredentialsId,
					},
					Certificate: objWithId{
						Id: buildCredential.CertificateId,
					},
					ProvisioningProfile: objWithId{
						Id: buildCredential.ProvisioningProfileId,
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
