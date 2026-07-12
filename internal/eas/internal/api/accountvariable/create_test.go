package accountvariable

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestCreate(t *testing.T) {
	expectedAccountId := "test-account-id"
	expectedData := &Data{
		Id:           "test-id",
		Name:         "test-name",
		Value:        "test-value",
		Visibility:   "PUBLIC",
		Environments: []string{"preview", "production"},
		CreatedAt:    "2024-01-01T00:00:00.000Z",
		UpdatedAt:    "2024-01-01T00:00:00.000Z",
	}

	expectedVariables := map[string]any{
		"accountId":    expectedAccountId,
		"name":         expectedData.Name,
		"value":        expectedData.Value,
		"environments": expectedData.Environments,
		"visibility":   expectedData.Visibility,
	}

	input := &CreateData{
		AccountId:    expectedAccountId,
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
