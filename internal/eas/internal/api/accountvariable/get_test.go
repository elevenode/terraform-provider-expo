package accountvariable

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

var expectedAccountId = "test-account-id"

var expectedData = &Data{
	Id:           "test-id",
	Name:         "test-name",
	Value:        "test-value",
	Visibility:   "PUBLIC",
	Environments: []string{"development"},
	CreatedAt:    "2024-01-01T00:00:00.000Z",
	UpdatedAt:    "2024-01-01T00:00:00.000Z",
}

var mockResponse = getResponse{Account: accountQuery{ById: accountById{Data: []Data{*expectedData}}}}

func TestGetById(t *testing.T) {
	getData := &GetData{
		AccountId: expectedAccountId,
		Id:        expectedData.Id,
	}
	expectedVariables := map[string]any{"accountId": expectedAccountId}

	config := utils.TestConfig[GetData, Data, getResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Get",
		Input:              getData,
		MockResponse:       mockResponse,
		ExpectedQuery:      getQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
