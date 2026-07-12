package appidentifier

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestCreate(t *testing.T) {
	identifier := "test-identifier"
	accountId := "test-account-id"
	expectedData := &Data{
		Id:         "test-id",
		Identifier: identifier,
	}

	expectedVariables := map[string]any{
		"identifier": identifier,
		"accountId":  accountId,
	}

	mockResponse := createResponse{
		CreateAppIdentifier: createAppBundleIdentifier{Data: *expectedData},
	}

	input := &CreateData{
		AccountId:  accountId,
		Identifier: identifier,
	}

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
