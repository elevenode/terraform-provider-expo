package appcredentials

import "github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/appbuildcredentials"

func mapData(data data) Data {
	buildCredentials := []appbuildcredentials.Data{}

	for _, b := range data.BuildCredentials {
		buildCredential := appbuildcredentials.Data{
			Id:                    b.Id,
			DistributionType:      b.DistributionType,
			AppCredentialsId:      b.AppCredentials.Id,
			ProvisioningProfileId: b.ProvisioningProfile.Id,
			CertificateId:         b.Certificate.Id,
		}
		buildCredentials = append(buildCredentials, buildCredential)
	}

	var appStoreApiKeyId *string = nil
	if data.AppStoreApiKey != nil {
		appStoreApiKeyId = &data.AppStoreApiKey.Id
	}

	var pushKeyId *string = nil
	if data.PushKey != nil {
		pushKeyId = &data.PushKey.Id
	}

	return Data{
		Id:               data.Id,
		AppId:            data.App.Id,
		AppIdentifierId:  data.AppIdentifier.Id,
		AppStoreApiKeyId: appStoreApiKeyId,
		PushKeyId:        pushKeyId,
		BuildCredentials: buildCredentials,
	}
}
