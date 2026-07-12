package team

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestUpdate(t *testing.T) {
	expectedData := &Data{
		Id:         "test-id",
		Name:       "test-name",
		Identifier: "test-identifier",
		Type:       "test-type",
		AccountId:  "test-account-id",
	}

	expectedVariables := map[string]any{
		"id":   expectedData.Id,
		"name": expectedData.Name,
		"type": expectedData.Type,
	}

	input := &UpdateData{
		Id:   expectedData.Id,
		Name: expectedData.Name,
		Type: expectedData.Type,
	}

	mockResponse := updateTeamResponse{
		UpdateAppleTeam: updateTeam{
			Data: &data{
				Id:         expectedData.Id,
				Name:       expectedData.Name,
				Identifier: expectedData.Identifier,
				Type:       expectedData.Type,
				Account:    account{Id: expectedData.AccountId},
			},
		},
	}

	config := utils.TestConfig[UpdateData, Data, updateTeamResponse, Service]{
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
