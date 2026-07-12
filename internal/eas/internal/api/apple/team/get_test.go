package team

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/graphql"
	"github.com/stretchr/testify/assert"
)

func TestGetByIdentifier(t *testing.T) {
	input := GetByIdentifierData{
		Identifier: "test-identifier",
		AccountId:  "test-account-id",
	}

	expectedData := &Data{
		Id:         "test-id",
		Name:       "test-name",
		Identifier: input.Identifier,
		Type:       "test-type",
		AccountId:  input.AccountId,
	}

	expectedVariables := map[string]any{
		"identifier": expectedData.Identifier,
		"accountId":  expectedData.AccountId,
	}

	mockData := &data{
		Id:         expectedData.Id,
		Name:       expectedData.Name,
		Identifier: expectedData.Identifier,
		Type:       expectedData.Type,
		Account:    account{Id: expectedData.AccountId},
	}
	mockResponse := getTeamResponse{GetTeam: getTeam{Data: mockData}}

	graphQLMock := graphql.NewGraphQLMock(mockResponse)

	service := NewService(graphQLMock)

	actualData, actualErr := service.GetByIdentifier(input)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, getQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
