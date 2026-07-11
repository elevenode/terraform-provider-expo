package appvariable

import (
	"testing"

	"terraform-provider-eas/internal/eas/internal/utils"
)

func TestCreate(t *testing.T) {
	expectedAppId := "test-app-id"
	expectedData := &Data{
		Id:           "test-id",
		Name:         "test-name",
		Value:        "test-value",
		Visibility:   "PUBLIC",
		Environments: []string{"preview", "production"},
	}

	expectedVariables := map[string]any{
		"appId":        expectedAppId,
		"name":         expectedData.Name,
		"value":        expectedData.Value,
		"environments": expectedData.Environments,
		"visibility":   expectedData.Visibility,
	}

	input := &CreateData{
		AppId:        expectedAppId,
		Name:         expectedData.Name,
		Value:        expectedData.Value,
		Visibility:   expectedData.Visibility,
		Environments: expectedData.Environments,
	}

	mockResponse := createResponse{EnvironmentVariable: createEnvironmentVariable{Data: expectedData}}

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
