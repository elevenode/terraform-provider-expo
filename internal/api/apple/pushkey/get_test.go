package pushkey

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/utils"
)

func TestGetByIdentifier(t *testing.T) {
	getData := &GeyByIdentifierData{
		Identifier: "test-identifier",
		AccountId:  "test-account-id",
	}
	expectedData := &Data{
		Id:         "test-id",
		Identifier: getData.Identifier,
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
