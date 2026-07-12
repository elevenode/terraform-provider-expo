package androidtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/elevenode/terraform-provider-expo/internal/eas/test/utils"
)

func TestCreateAndDeleteKeystore(t *testing.T) {
	input := eas.CreateAndroidKeystoreData{
		AccountId:        utils.AccountId,
		KeystoreBase64:   utils.KeystoreBase64,
		KeyAlias:         "test-key-alias",
		KeyPassword:      "test-key-password",
		KeystorePassword: "test-keystore-password",
	}

	actualData, actualErr := utils.Client.Android.Keystore.Create(input)

	assert.NoError(t, actualErr)
	assert.Equal(t, input.KeyAlias, actualData.KeyAlias)
	// assert.Equal(t, input.KeyPassword, actualData.KeyPassword) TODO EAS is not returning the key password
	assert.Equal(t, input.KeystoreBase64, actualData.Keystore)
	assert.Equal(t, input.KeystorePassword, actualData.KeystorePassword)

	_, actualErr = utils.Client.Android.Keystore.Delete(actualData.Id)
	assert.Equal(t, nil, actualErr)
}
