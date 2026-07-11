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
		Id:                    input.Id,
		DistributionType:      "APP_STORE",
		CertificateId:         "test-certificate-id",
		ProvisioningProfileId: "test-provisoning-profile-id",
		AppCredentialsId:      input.AppCredentialsId,
	}

	expectedVariables := map[string]any{"appId": input.AppId}

	mockResponse := getResponse{
		AppByAppId: appByAppId{
			IosAppCredentials: []iosAppCredentials{{
				Id: expectedData.AppCredentialsId,
				Data: []data{{
					Id:               expectedData.Id,
					DistributionType: expectedData.DistributionType,
					ProvisioningProfile: objWithId{
						Id: expectedData.ProvisioningProfileId,
					},
					Certificate: objWithId{
						Id: expectedData.CertificateId,
					},
					AppCredentials: objWithId{
						Id: expectedData.AppCredentialsId,
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
