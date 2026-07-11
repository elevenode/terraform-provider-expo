package appcredentials

import "terraform-provider-eas/internal/eas/internal/api/android/appbuildcredentials"

func mapData(data data) Data {
	buildCredentials := []appbuildcredentials.Data{}
	for _, b := range data.BuildCredentials {
		buildCredential := appbuildcredentials.Data{
			Id:               b.Id,
			Name:             b.Name,
			AppCredentialsId: data.Id,
			KeystoreId:       b.Keystore.Id,
		}
		buildCredentials = append(buildCredentials, buildCredential)
	}
	output := Data{
		Id:                        data.Id,
		AppId:                     data.App.Id,
		Identifier:                data.Identifier,
		GoogleServiceAccountKeyId: data.GoogleServiceAccountKey.Id,
		BuildCredentials:          buildCredentials,
	}
	if data.FCMKey != nil {
		output.FCMKey = &data.FCMKey.KeyJson
		output.FCMKeyId = &data.FCMKey.Id
	}
	return output
}
