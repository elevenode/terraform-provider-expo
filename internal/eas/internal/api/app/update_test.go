package app

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestUpdate(t *testing.T) {
	expectedData := &Data{
		Id:   "test-id",
		Name: "test-name",
	}

	expectedVariables := map[string]any{
		"id":   expectedData.Id,
		"name": expectedData.Name,
	}

	input := &UpdateData{
		Id:   expectedData.Id,
		Name: expectedData.Name,
	}

	mockResponse := updateResponse{UpdateApp: updateApp{Data: expectedData}}
	config := utils.TestConfig[UpdateData, Data, updateResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Update",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      updateQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
