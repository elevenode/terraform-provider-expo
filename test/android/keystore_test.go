package androidtest

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/fintreal/eas-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateKeystore(t *testing.T) {
	input := eas.CreateAndroidKeystoreData{
		AccountId:        utils.AccountId,
		KeystoreBase64:   "test-keystore-base64",
		KeyAlias:         "test-key-alias",
		KeyPassword:      "test-key-password",
		KeystorePassword: "test-keystore-password",
	}

	actualData, actualErr := utils.Client.Android.Keystore.Create(input)
	assert.Equal(t, input.KeyAlias, actualData.KeyAlias)
	assert.Equal(t, input.KeyPassword, actualData.KeyPassword)
	assert.Equal(t, input.KeystorePassword, actualData.KeystorePassword)
	assert.Equal(t, nil, actualErr)
}

func TestCreateAndDeleteKeystore(t *testing.T) {
	input := eas.CreateAndroidKeystoreData{
		AccountId:        utils.AccountId,
		KeystoreBase64:   "test-keystore-base64",
		KeyAlias:         "test-key-alias",
		KeyPassword:      "test-key-password",
		KeystorePassword: "test-keystore-password",
	}

	actualData, actualErr := utils.Client.Android.Keystore.Create(input)
	assert.Equal(t, input.KeyAlias, actualData.KeyAlias)
	assert.Equal(t, input.KeyPassword, actualData.KeyPassword)
	assert.Equal(t, input.KeystorePassword, actualData.KeystorePassword)
	assert.Equal(t, nil, actualErr)

	_, actualErr = utils.Client.Android.Keystore.Delete(actualData.Id)
	assert.Equal(t, nil, actualErr)
}
