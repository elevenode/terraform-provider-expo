package androidtest

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/elevenode/terraform-provider-expo/internal/eas/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAppBuildCredentials(t *testing.T) {
	input := eas.GetAndroidAppBuildCredentialsData{
		Id:               utils.ImmutableAndroidAppBuildCredentialsId,
		AppId:            utils.ImmutableAppId,
		AppCredentialsId: utils.ImmutableAndroidAppCredentialsId,
	}
	actualData, actualErr := utils.Client.Android.AppBuildCredentials.Get(input)

	expectedData := &eas.AndroidAppBuildCredentialsData{
		Id:               input.Id,
		Name:             "Default",
		KeystoreId:       utils.ImmutableKeystoreId,
		AppCredentialsId: input.AppCredentialsId,
	}

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestCreateAndDeleteAppBuildCredentials(t *testing.T) {
	input := eas.CreateAndroidAppBuildCredentialsData{
		AppCredentialsId: utils.ImmutableAndroidAppCredentialsId,
		KeystoreId:       utils.ImmutableKeystoreId,
		Name:             utils.GenerateRandomString(10),
	}

	actualData, actualErr := utils.Client.Android.AppBuildCredentials.Create(input)

	assert.Equal(t, input.AppCredentialsId, actualData.AppCredentialsId)
	assert.Equal(t, input.KeystoreId, actualData.KeystoreId)
	assert.Equal(t, input.Name, actualData.Name)
	assert.Equal(t, nil, actualErr)

	_, actualErr = utils.Client.Android.AppBuildCredentials.Delete(actualData.Id)

	assert.Equal(t, nil, actualErr)

}
