package environmentvariable

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/utils"
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
