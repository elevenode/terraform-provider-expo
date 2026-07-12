package appletest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/elevenode/terraform-provider-expo/internal/eas/test/utils"
)

func TestGetAppleTeamByIdentifier(t *testing.T) {
	input := &eas.GetByIdentifierAppleTeamData{
		Identifier: "TEST_IDENTIFIER",
		AccountId:  utils.AccountId,
	}

	actualData, actualErr := utils.Client.Apple.Team.GetByIdentifier(*input)

	expectedData := &eas.AppleTeamData{
		Id:         utils.ImmutableAppleTeamId,
		AccountId:  utils.AccountId,
		Identifier: "TEST_IDENTIFIER",
		Name:       "TEST_NAME",
		Type:       "COMPANY_OR_ORGANIZATION",
	}

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestCreateAppleTeamByIdentifier(t *testing.T) {
	expectedData := &eas.AppleTeamData{
		AccountId:  utils.AccountId,
		Identifier: utils.GenerateRandomString(10),
		Name:       utils.GenerateRandomString(10),
		Type:       "COMPANY_OR_ORGANIZATION", // RANDOMIZE
	}

	input := eas.CreateAppleTeamData{
		Identifier: expectedData.Identifier,
		Name:       expectedData.Name,
		Type:       expectedData.Type,
		AccountId:  utils.AccountId,
	}

	actualData, actualErr := utils.Client.Apple.Team.Create(input)

	assert.Equal(t, expectedData.AccountId, actualData.AccountId)
	assert.Equal(t, expectedData.Identifier, actualData.Identifier)
	assert.Equal(t, expectedData.Name, actualData.Name)
	assert.Equal(t, expectedData.Type, actualData.Type)
	assert.Equal(t, nil, actualErr)
}

func TestUpdateAppleTeamByIdentifier(t *testing.T) {
	expectedData := &eas.AppleTeamData{
		Id:         utils.MutableAppleTeamId,
		AccountId:  utils.AccountId,
		Identifier: "TEST_UPDATE_TEAM",
		Name:       utils.GenerateRandomString(10),
		Type:       "COMPANY_OR_ORGANIZATION", // RANDOMIZE
	}

	input := eas.UpdateAppleTeamData{
		Id:   expectedData.Id,
		Name: expectedData.Name,
		Type: expectedData.Type,
	}

	actualData, actualErr := utils.Client.Apple.Team.Update(input)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
