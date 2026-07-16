package updatechannel

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestUpdate(t *testing.T) {
	expectedData := &Data{
		Id:            "test-id",
		Name:          "test-name",
		BranchMapping: `{"version":0,"data":[{"branchId":"test-branch-id","branchMappingLogic":"true"}]}`,
	}
	input := &UpdateData{
		Id:            expectedData.Id,
		BranchMapping: expectedData.BranchMapping,
	}
	expectedVariables := map[string]any{
		"channelId":     expectedData.Id,
		"branchMapping": expectedData.BranchMapping,
	}

	mockResponse := updateResponse{UpdateChannel: editUpdateChannel{Data: expectedData}}

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
