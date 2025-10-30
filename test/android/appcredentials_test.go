package androidtest

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/fintreal/eas-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAppCredentials(t *testing.T) {
	input := eas.GetAndroidAppCredentialsData{
		Id:    "9994bb7d-e1d5-4d50-8429-addfa2f24f72",
		AppId: utils.ImmutableAppId,
	}
	actualData, actualErr := utils.Client.Android.AppCredentials.Get(input)

	fcmKeyId := "c5986341-d263-4b1b-a1db-56f0b7d766fa"

	expectedData := &eas.AndroidAppCredentialsData{
		Id:                        input.Id,
		AppId:                     input.AppId,
		Identifier:                "my.test.app.identifier",
		GoogleServiceAccountKeyId: "36b45ce5-1cf3-4e29-a04d-88fb3c4b5683",
		FCMKeyId:                  &fcmKeyId,
		FCMKey:                    &utils.FCMKey,
		BuildCredentials: []eas.AndroidAppBuildCredentialsData{{
			AppCredentialsId: input.Id,
			Id:               utils.ImmutableAndroidAppBuildCredentialsId,
			Name:             "Default",
			KeystoreId:       utils.ImmutableKeystoreId,
		}},
	}
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedData, actualData)
}

func TestCreateAndDeleteAppCredentials(t *testing.T) {
	input := eas.CreateAndroidAppCredentialsData{
		AppId:                     utils.ImmutableAppId,
		Identifier:                utils.GenerateRandomString(10),
		GoogleServiceAccountKeyId: "36b45ce5-1cf3-4e29-a04d-88fb3c4b5683",
	}

	actualData, actualErr := utils.Client.Android.AppCredentials.Create(input)
	assert.Equal(t, input.AppId, actualData.AppId)
	assert.Equal(t, input.Identifier, actualData.Identifier)
	assert.Equal(t, input.GoogleServiceAccountKeyId, actualData.GoogleServiceAccountKeyId)
	assert.Equal(t, nil, actualErr)

	_, actualErr = utils.Client.Android.AppCredentials.Delete(actualData.Id)

	assert.Equal(t, nil, actualErr)
}
