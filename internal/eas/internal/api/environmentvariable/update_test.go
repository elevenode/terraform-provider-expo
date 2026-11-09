package environmentvariable

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestUpdate(t *testing.T) {
	expectedData := &Data{
		Id:           "test-id",
		Name:         "test-name-updated",
		Value:        "test-value-updated",
		Visibility:   "SENSITIVE",
		Environments: []string{"development", "preview", "production"},
		CreatedAt:    "2024-01-01T00:00:00.000Z",
		UpdatedAt:    "2024-01-02T00:00:00.000Z",
	}

	expectedVariables := map[string]any{
		"id":           expectedData.Id,
		"name":         expectedData.Name,
		"value":        expectedData.Value,
		"environments": expectedData.Environments,
		"visibility":   expectedData.Visibility,
	}

	input := &UpdateData{
		Id:           expectedData.Id,
		Name:         expectedData.Name,
		Value:        expectedData.Value,
		Visibility:   expectedData.Visibility,
		Environments: expectedData.Environments,
	}

	mockResponse := updateResponse{EnvironmentVariable: updateEnvironmentVariable{Data: expectedData}}

	config := utils.TestConfig[UpdateData, Data, updateResponse, BaseService]{
		NewServiceFunction: NewBaseService,
		FunctionUnderTest:  "Update",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      updateQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}

func TestUpdateFileVariable(t *testing.T) {
	expectedData := &Data{
		Id:           "test-id",
		Name:         "GOOGLE_SERVICES_JSON",
		Value:        "eyJwcm9qZWN0X2luZm8iOnt9fQ==",
		Visibility:   "SECRET",
		Environments: []string{"production"},
		Type:         "FILE_BASE64",
		FileName:     "google-services.json",
	}

	expectedVariables := map[string]any{
		"id":           expectedData.Id,
		"name":         expectedData.Name,
		"value":        expectedData.Value,
		"environments": expectedData.Environments,
		"visibility":   expectedData.Visibility,
		"type":         expectedData.Type,
		"fileName":     expectedData.FileName,
	}

	input := &UpdateData{
		Id:           expectedData.Id,
		Name:         expectedData.Name,
		Value:        expectedData.Value,
		Visibility:   expectedData.Visibility,
		Environments: expectedData.Environments,
		Type:         expectedData.Type,
		FileName:     expectedData.FileName,
	}

	mockResponse := updateResponse{EnvironmentVariable: updateEnvironmentVariable{Data: expectedData}}

	config := utils.TestConfig[UpdateData, Data, updateResponse, BaseService]{
		NewServiceFunction: NewBaseService,
		FunctionUnderTest:  "Update",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      updateQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
