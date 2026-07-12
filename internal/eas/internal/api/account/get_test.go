package account

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestGetByName(t *testing.T) {
	expectedData := &Data{
		Id:   "test-account-id",
		Name: "test-account-name",
	}
	expectedVariables := map[string]any{"name": expectedData.Name}

	mockResponse := getResponse{Account: account{ByName: expectedData}}

	config := utils.TestConfig[string, Data, getResponse, AccountService]{
		NewServiceFunction: NewAccountService,
		FunctionUnderTest:  "GetByName",
		Input:              &expectedData.Name,
		MockResponse:       mockResponse,
		ExpectedQuery:      getByNameQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
