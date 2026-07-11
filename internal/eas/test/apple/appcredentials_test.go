package appletest

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/fintreal/eas-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAppCredentials(t *testing.T) {
	input := eas.GetAppCredentialsData{
		Id:    utils.ImmutableAppleAppCredentialsId,
		AppId: utils.ImmutableAppId,
	}
	actualData, actualErr := utils.Client.Apple.AppCredentials.Get(input)

	expectedData := &eas.AppCredentialsData{
		Id:               input.Id,
		AppId:            input.AppId,
		AppIdentifierId:  "90289e95-99c2-4b21-88cc-8d84b6ab7477",
		AppStoreApiKeyId: &utils.ImmutableAppStoreApiKeyId,
		BuildCredentials: []eas.AppBuildCredentialsData{
			{
				Id:                    "6ed2748a-f54c-4da1-8187-16593ef195cd",
				DistributionType:      "APP_STORE",
				CertificateId:         "702635c5-3aa1-477c-83b6-bb66a1644aad",
				ProvisioningProfileId: "72157c17-10db-4851-8633-afd5a08384ce",
				AppCredentialsId:      "11f2b3f8-ddad-4626-8984-2b96efb28d3c",
			},
			{
				Id:                    "bc23563b-b905-481b-941e-8ae412b994b3",
				DistributionType:      "DEVELOPMENT",
				CertificateId:         "702635c5-3aa1-477c-83b6-bb66a1644aad",
				ProvisioningProfileId: "72157c17-10db-4851-8633-afd5a08384ce",
				AppCredentialsId:      "11f2b3f8-ddad-4626-8984-2b96efb28d3c",
			},
		},
	}

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestCreateUpdateAndDeleteAppCredentials(t *testing.T) {
	input := eas.CreateAppCredentialsData{
		AppId:           utils.ImmutableAppId,
		AppIdentifierId: utils.MutableAppIdentifierId,
	}

	actualData, actualErr := utils.Client.Apple.AppCredentials.Create(input)
	assert.Equal(t, input.AppId, actualData.AppId)
	assert.Equal(t, input.AppIdentifierId, actualData.AppIdentifierId)
	assert.Nil(t, actualData.AppStoreApiKeyId)
	assert.Nil(t, actualData.PushKeyId)
	assert.Equal(t, nil, actualErr)

	updateInput := eas.UpdateAppCredentialsData{
		Id:               actualData.Id,
		AppStoreApiKeyId: &utils.ImmutableAppStoreApiKeyId,
		PushKeyId:        nil,
	}

	actualData, actualErr = utils.Client.Apple.AppCredentials.Update(updateInput)

	assert.Equal(t, input.AppId, actualData.AppId)
	assert.Equal(t, input.AppIdentifierId, actualData.AppIdentifierId)
	assert.Equal(t, utils.ImmutableAppStoreApiKeyId, *actualData.AppStoreApiKeyId)
	assert.Nil(t, actualData.PushKeyId)
	assert.Equal(t, nil, actualErr)

	_, actualErr = utils.Client.Apple.AppCredentials.Delete(actualData.Id)

	assert.Equal(t, nil, actualErr)
}
