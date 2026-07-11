package googleserviceaccountkey

import (
	"testing"

	"terraform-provider-eas/internal/eas/internal/utils"
)

func TestGetByProjectIdentifier(t *testing.T) {
	getData := &GetByProjectIdentifierData{
		ProjectIdentifier: "test-identifier",
		AccountId:         "test-account-id",
	}
	expectedData := &Data{
		Id:                "test-id",
		ProjectIdentifier: getData.ProjectIdentifier,
		ClientEmail:       "test-client-email",
		ClientIdentifier:  "test-client-identifier",
	}
	expectedVariables := map[string]any{"accountId": getData.AccountId}

	mockResponse := getResponse{
		Account: utils.Account[getAccount]{
			ById: getAccount{
				Data: []Data{*expectedData},
			},
		},
	}

	config := utils.TestConfig[GetByProjectIdentifierData, Data, utils.AccountResponse[getAccount], Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "GetByProjectIdentifier",
		Input:              getData,
		MockResponse:       mockResponse,
		ExpectedQuery:      getQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
