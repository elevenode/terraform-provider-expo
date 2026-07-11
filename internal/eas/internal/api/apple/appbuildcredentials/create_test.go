package appbuildcredentials

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/utils"
)

func TestCreate(t *testing.T) {
	expectedData := &Data{
		Id:                    "test-id",
		DistributionType:      "APP_STORE",
		ProvisioningProfileId: "test-provisioning-profile-id",
		CertificateId:         "test-certificate-id",
		AppCredentialsId:      "test-app-credentials-id",
	}

	expectedVariables := map[string]any{
		"distributionType":      expectedData.DistributionType,
		"certificateId":         expectedData.CertificateId,
		"provisioningProfileId": expectedData.ProvisioningProfileId,
		"appCredentialsId":      expectedData.AppCredentialsId,
	}

	mockResponse := createResponse{
		IosAppBuildCredentials: createIosAppBuildCredentials{
			Data: data{
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
			},
		},
	}

	input := &CreateData{
		CertificateId:         expectedData.CertificateId,
		DistributionType:      expectedData.DistributionType,
		ProvisioningProfileId: expectedData.ProvisioningProfileId,
		AppCredentialsId:      expectedData.AppCredentialsId,
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
