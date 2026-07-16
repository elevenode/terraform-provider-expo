package updatebranch

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestGetByName(t *testing.T) {
	appId := "test-app-id"
	expectedData := &Data{
		Id:   "test-id",
		Name: "test-name",
	}
	input := &GetByNameData{
		AppId: appId,
		Name:  expectedData.Name,
	}
	expectedVariables := map[string]any{
		"appId": appId,
		"name":  expectedData.Name,
	}

	mockResponse := getResponse{App: appQuery{ById: branchByAppId{Data: expectedData}}}

	config := utils.TestConfig[GetByNameData, Data, getResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "GetByName",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      getByNameQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
