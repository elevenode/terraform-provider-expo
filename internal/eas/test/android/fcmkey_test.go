package androidtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/elevenode/terraform-provider-expo/internal/eas/test/utils"
)

func TestCreateAppFCMKey(t *testing.T) {
	input := eas.CreateFCMKey{
		AccountId:        utils.AccountId,
		AppCredentialsId: "158aee51-3df6-45ba-9b18-04f33ca0cd39",
		KeyJson:          utils.FCMKey,
	}

	actualData, actualErr := utils.Client.Android.FCMKey.Create(input)
	assert.JSONEq(t, input.KeyJson, actualData.KeyJson)
	assert.NoError(t, actualErr)
}
