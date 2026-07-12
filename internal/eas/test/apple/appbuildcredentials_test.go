package appletest

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/elevenode/terraform-provider-expo/internal/eas/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAppBuildCredentials(t *testing.T) {
	input := eas.GetAppBuildCredentialsData{
		Id:               "6ed2748a-f54c-4da1-8187-16593ef195cd",
		AppId:            utils.ImmutableAppId,
		AppCredentialsId: utils.ImmutableAppleAppCredentialsId,
	}
	actualData, actualErr := utils.Client.Apple.AppBuildCredentials.Get(input)

	expectedData := &eas.AppBuildCredentialsData{
		Id:                    input.Id,
		DistributionType:      "APP_STORE",
		CertificateId:         utils.ImmutableCertificateId,
		ProvisioningProfileId: "72157c17-10db-4851-8633-afd5a08384ce", // TODO refactor
		AppCredentialsId:      utils.ImmutableAppleAppCredentialsId,
	}

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestCreateAndDeleteAppBuildCredentials(t *testing.T) {
	input := eas.CreateAppBuildCredentialsData{
		DistributionType:      "APP_STORE",
		ProvisioningProfileId: utils.ImmutableProvisioningProfileId,
		CertificateId:         utils.ImmutableCertificateId,
		AppCredentialsId:      utils.MutableAppCredentialsId,
	}

	actualData, actualErr := utils.Client.Apple.AppBuildCredentials.Create(input)

	assert.Equal(t, input.DistributionType, actualData.DistributionType)
	assert.Equal(t, input.CertificateId, actualData.CertificateId)
	assert.Equal(t, input.AppCredentialsId, actualData.AppCredentialsId)
	assert.Equal(t, input.ProvisioningProfileId, actualData.ProvisioningProfileId)
	assert.Equal(t, nil, actualErr)

	_, actualErr = utils.Client.Apple.AppBuildCredentials.Delete(actualData.Id)

	assert.Equal(t, nil, actualErr)

}
