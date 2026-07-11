package me

import (
	"testing"

	"terraform-provider-eas/internal/eas/internal/utils"
)

func TestGet(t *testing.T) {
	var expectedData = &Data{
		Id:          "test-id",
		DisplayName: "Test Display Name",
	}
	mockResponse := getResponse{Data: expectedData}
	config := utils.TestConfig[any, Data, getResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Get",
		Input:              nil,
		MockResponse:       mockResponse,
		ExpectedQuery:      getQuery,
		ExpectedVariables:  map[string]any{},
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
