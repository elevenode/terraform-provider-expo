package appletest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"terraform-provider-eas/internal/eas"
	"terraform-provider-eas/internal/eas/test/utils"
)

func TestGetAppStoreApiKeyByIdentifier(t *testing.T) {
	input := eas.GetByIdentifierAppStoreApiKeyData{
		Identifier: "349P6U74M8",
		AccountId:  utils.AccountId,
	}

	actualData, actualErr := utils.Client.Apple.AppStoreApiKey.GetByIdentifier(input)

	expectedData := &eas.AppStoreApiKeyData{
		Id:               "564e9d75-ff77-4860-92ee-7c0ab2066c82",
		Identifier:       input.Identifier,
		IssuerIdentifier: "c226a277-e124-405a-a571-794533cd587a",
		Name:             "App Store API Key",
	}

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
