package team

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestCreate(t *testing.T) {
	expectedData := &Data{
		Id:         "test-id",
		Name:       "test-name",
		Identifier: "test-identifier",
		Type:       "test-type",
		AccountId:  "test-account-id",
	}

	expectedVariables := map[string]any{
		"identifier": expectedData.Identifier,
		"name":       expectedData.Name,
		"accountId":  expectedData.AccountId,
		"type":       expectedData.Type,
	}

	input := &CreateData{
		Name:       expectedData.Name,
		Identifier: expectedData.Identifier,
		Type:       expectedData.Type,
		AccountId:  expectedData.AccountId,
	}

	mockResponse := createTeamResponse{
		CreateTeam: createTeam{
			Data: &data{
				Id:         expectedData.Id,
				Name:       expectedData.Name,
				Identifier: expectedData.Identifier,
				Type:       expectedData.Type,
				Account:    account{Id: expectedData.AccountId},
			},
		},
	}

	config := utils.TestConfig[CreateData, Data, createTeamResponse, Service]{
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
