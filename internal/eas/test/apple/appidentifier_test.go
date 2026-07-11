package appletest

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/fintreal/eas-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
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
