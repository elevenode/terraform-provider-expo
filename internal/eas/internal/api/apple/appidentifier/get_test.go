package appidentifier

import (
	"testing"

	"terraform-provider-eas/internal/eas/internal/utils"
)

func TestGetByIdentifier(t *testing.T) {
	identifier := "test-identifier"
	accountId := "test-account-id"
	input := &GetByIdentifierData{
		Identifier: identifier,
		AccountId:  accountId,
	}
	expectedData := &Data{
		Id:         "test-id",
		Identifier: identifier,
	}

	expectedVariables := map[string]any{
		"identifier": identifier,
		"accountId":  accountId,
	}
	mockResponse := getMockResponse(expectedData)

	config := utils.TestConfig[GetByIdentifierData, Data, utils.AccountResponse[getResponse], Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "GetByIdentifier",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      getQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}

func getMockResponse(input *Data) utils.AccountResponse[getResponse] {
	mockData := Data{
		Id:         input.Id,
		Identifier: input.Identifier,
	}
	return utils.AccountResponse[getResponse]{
		Account: utils.Account[getResponse]{
			ById: getResponse{
				Data: []Data{mockData},
			},
		},
	}
}
