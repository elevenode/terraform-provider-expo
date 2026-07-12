package appletest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/elevenode/terraform-provider-expo/internal/eas/test/utils"
)

func TestGetAppleAppIdentifierByIdentifier(t *testing.T) {
	input := eas.GetByIdentifierAppleAppIdentifierData{
		Identifier: utils.ImmutableAppIdentifierName,
		AccountId:  utils.AccountId,
	}

	actualData, actualErr := utils.Client.Apple.AppIdentifier.GetByIdentifier(input)

	expectedData := &eas.AppleAppIdentifierData{
		Id:         utils.ImmutableAppIdentifierId,
		Identifier: input.Identifier,
	}

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestCreateAppleAppIdentifier(t *testing.T) {
	expectedIdentifier := utils.GenerateRandomString(10)

	input := eas.CreateAppleAppIdentifierData{
		AccountId:  utils.AccountId,
		Identifier: expectedIdentifier,
	}

	actualData, actualErr := utils.Client.Apple.AppIdentifier.Create(input)

	assert.Equal(t, expectedIdentifier, actualData.Identifier)
	assert.Equal(t, nil, actualErr)
}
