package appvariable

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

var expectedAppId = "test-app-id"

var expectedData = &Data{
	Id:           "test-id",
	Name:         "test-name",
	Value:        "test-value",
	Visibility:   "PUBLIC",
	Environments: []string{"development"},
}

var mockResponse = getResponse{AppByAppId: appById{Data: []Data{*expectedData}}}

func TestGetByName(t *testing.T) {
	getData := &GetByNameData{
		AppId: expectedAppId,
		Name:  expectedData.Name,
	}
	expectedVariables := map[string]any{"appId": expectedAppId}

	config := utils.TestConfig[GetByNameData, Data, getResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "GetByName",
		Input:              getData,
		MockResponse:       mockResponse,
		ExpectedQuery:      getQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}

func TestGetById(t *testing.T) {
	getData := &GetData{
		AppId: expectedAppId,
		Id:    expectedData.Id,
	}
	expectedVariables := map[string]any{"appId": expectedAppId}

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
