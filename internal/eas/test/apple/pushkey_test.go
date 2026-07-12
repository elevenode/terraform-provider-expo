package appletest

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/elevenode/terraform-provider-expo/internal/eas/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetPushKeyByIdentifier(t *testing.T) {
	input := eas.GetByIdentifierPushKeyData{
		Identifier: "M6SPYT2C2L",
		AccountId:  utils.AccountId,
	}

	actualData, actualErr := utils.Client.Apple.PushKey.GetByIdentifier(input)

	expectedData := &eas.PushKeyData{
		Id:         "be4f6cfa-0f92-4db5-b875-d2a1973e8ad7",
		Identifier: input.Identifier,
	}

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
