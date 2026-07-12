package app

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestGet(t *testing.T) {
	expectedData := &Data{
		Id:   "test-id",
		Name: "test-name",
		Slug: "test-slug",
	}

	config := utils.TestConfig[string, Data, getResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Get",
		Input:              &expectedData.Id,
		MockResponse:       getResponse{Data: expectedData},
		ExpectedQuery:      getQuery,
		ExpectedVariables:  map[string]any{"id": expectedData.Id},
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}

func TestGetByFullName(t *testing.T) {
	expectedData := &Data{
		Id:   "test-id",
		Name: "test-name",
		Slug: "test-slug",
	}

	config := utils.TestConfig[string, Data, getByFullNameResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "GetByFullName",
		Input:              &expectedData.Id,
		MockResponse:       getByFullNameResponse{App: &app{Data: expectedData}},
		ExpectedQuery:      getByFullNameQuery,
		ExpectedVariables:  map[string]any{"fullName": expectedData.Id},
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
