package appvariable

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
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
		"type":         "STRING",
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

func TestCreateFileVariable(t *testing.T) {
	expectedAppId := "test-app-id"
	expectedData := &Data{
		Id:           "test-id",
		Name:         "GOOGLE_SERVICES_JSON",
		Value:        "eyJwcm9qZWN0X2luZm8iOnt9fQ==",
		Visibility:   "SECRET",
		Environments: []string{"production"},
		Type:         TypeFileBase64,
		FileName:     "google-services.json",
	}
	expectedVariables := map[string]any{
		"appId":        expectedAppId,
		"name":         expectedData.Name,
		"value":        expectedData.Value,
		"environments": expectedData.Environments,
		"visibility":   expectedData.Visibility,
		"type":         TypeFileBase64,
		"fileName":     expectedData.FileName,
	}
	input := &CreateData{
		AppId:        expectedAppId,
		Name:         expectedData.Name,
		Value:        expectedData.Value,
		Visibility:   expectedData.Visibility,
		Environments: expectedData.Environments,
		Type:         TypeFileBase64,
		FileName:     expectedData.FileName,
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
