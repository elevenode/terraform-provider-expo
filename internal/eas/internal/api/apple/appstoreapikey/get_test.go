package appstoreapikey

import (
	"testing"

	"terraform-provider-eas/internal/eas/internal/utils"
)

func TestGetByIdentifier(t *testing.T) {
	getData := &GeyByIdentifierData{
		Identifier: "test-identifier",
		AccountId:  "test-account-id",
	}
	expectedData := &Data{
		Id:               "test-id",
		Name:             "test-name",
		IssuerIdentifier: "test-issuer-identifier",
		Identifier:       getData.Identifier,
	}
	expectedVariables := map[string]any{"accountId": getData.AccountId}

	mockResponse := getResponse{
		Account: utils.Account[getAccount]{
			ById: getAccount{
				Data: []Data{*expectedData},
			},
		},
	}

	config := utils.TestConfig[GeyByIdentifierData, Data, utils.AccountResponse[getAccount], Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "GetByIdentifier",
		Input:              getData,
		MockResponse:       mockResponse,
		ExpectedQuery:      getQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
