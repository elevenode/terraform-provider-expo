package updatebranch

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestCreate(t *testing.T) {
	appId := "test-app-id"
	expectedData := &Data{
		Id:   "test-id",
		Name: "test-name",
	}
	expectedVariables := map[string]any{
		"appId": appId,
		"name":  expectedData.Name,
	}
	input := &CreateData{
		AppId: appId,
		Name:  expectedData.Name,
	}

	mockResponse := createResponse{UpdateBranch: createUpdateBranch{Data: expectedData}}

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
