package app

import (
	"testing"

	"terraform-provider-eas/internal/eas/internal/utils"
)

func TestCreate(t *testing.T) {
	accountId := "test-account-id"
	expectedData := &Data{
		Id:   "test-id",
		Name: "test-name",
		Slug: "test-slug",
	}
	expectedVariables := map[string]any{
		"accountId": accountId,
		"name":      expectedData.Name,
		"slug":      expectedData.Slug,
	}
	input := &CreateData{
		AccountId: accountId,
		Name:      expectedData.Name,
		Slug:      expectedData.Slug,
	}

	mockResponse := createResponse{CreateApp: createApp{Data: expectedData}}

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
