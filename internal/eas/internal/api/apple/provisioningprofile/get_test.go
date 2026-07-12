package provisioningprofile

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestGet(t *testing.T) {
	input := &GetData{
		Id:        "test-id",
		AccountId: "test-account-id",
	}

	identifier := "test-identifier"

	expectedData := &Data{
		Id:                    input.Id,
		AppBundleIdentifierId: identifier,
		Base64:                "test-base64-string",
	}

	expectedVariables := map[string]any{"accountId": input.AccountId}

	mockResponse := getMockResponse(expectedData)

	config := utils.TestConfig[GetData, Data, utils.AccountResponse[getProvisioningProfilesResponse], Service]{
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

func getMockResponse(input *Data) utils.AccountResponse[getProvisioningProfilesResponse] {
	mockData := data{
		Id:                 input.Id,
		Base64:             input.Base64,
		AppleAppIdentifier: appleAppIdentifier{Id: input.AppBundleIdentifierId},
	}

	return utils.AccountResponse[getProvisioningProfilesResponse]{
		Account: utils.Account[getProvisioningProfilesResponse]{
			ById: getProvisioningProfilesResponse{
				Data: []data{mockData},
			},
		},
	}
}
