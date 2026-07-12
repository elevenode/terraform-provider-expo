package appletest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/elevenode/terraform-provider-expo/internal/eas/test/utils"
)

func TestGetProvisioningProfile(t *testing.T) {
	input := eas.GetProvisioningProfileData{
		Id:        utils.ImmutableProvisioningProfileId,
		AccountId: utils.AccountId,
	}

	actualData, actualErr := utils.Client.Apple.ProvisioningProfile.Get(input)

	expectedData := &eas.ProvisioningProfileData{
		Id:                    input.Id,
		Base64:                utils.ImmutableProvisioningProfileBase64,
		AppBundleIdentifierId: "554fe9ea-dadd-49ec-8300-1940661f352b",
	}
	assert.Equal(t, expectedData, actualData)
	assert.NoError(t, actualErr)
}

func TestCreateAndDeleteProvisioningProfile(t *testing.T) {
	expectedBase64 := utils.ImmutableProvisioningProfileBase64
	expectedAppBundleIdentifierId := "554fe9ea-dadd-49ec-8300-1940661f352b"

	input := eas.CreateProvisioningProfileData{
		Base64:                utils.ImmutableProvisioningProfileBase64,
		AppBundleIdentifierId: "554fe9ea-dadd-49ec-8300-1940661f352b",
		AccountId:             utils.AccountId,
	}

	actualData, actualErr := utils.Client.Apple.ProvisioningProfile.Create(input)
	assert.Equal(t, expectedBase64, actualData.Base64)
	assert.Equal(t, expectedAppBundleIdentifierId, actualData.AppBundleIdentifierId)
	assert.NoError(t, actualErr)

	_, actualErr = utils.Client.Apple.ProvisioningProfile.Delete(actualData.Id)
	assert.NoError(t, actualErr)
}
