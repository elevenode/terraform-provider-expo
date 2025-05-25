package androidtest

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/fintreal/eas-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateAppFCMKey(t *testing.T) {
	input := eas.CreateFCMKey{
		AccountId:        utils.AccountId,
		AppCredentialsId: "158aee51-3df6-45ba-9b18-04f33ca0cd39",
		KeyJson:          utils.FCMKey,
	}

	actualData, actualErr := utils.Client.Android.FCMKey.Create(input)
	assert.Equal(t, input.KeyJson, actualData.KeyJson)
	assert.Equal(t, nil, actualErr)
}
