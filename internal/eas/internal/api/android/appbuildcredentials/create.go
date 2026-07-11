package appbuildcredentials

type createAndroidAppBuildCredentials struct {
	Data data `json:"createAndroidAppBuildCredentials"`
}
type createResponse struct {
	AndroidAppBuildCredentials createAndroidAppBuildCredentials `json:"androidAppBuildCredentials"`
}

const createQuery = `
	mutation ($appCredentialsId: ID!, $keystoreId: ID!, $name: String!) {
    androidAppBuildCredentials {
      createAndroidAppBuildCredentials(
        androidAppCredentialsId: $appCredentialsId
        androidAppBuildCredentialsInput: { keystoreId: $keystoreId, isDefault: true, name: $name }
      ) {
        id
        name
        androidKeystore { id }
      }
    }
	}
`

func (service *service) Create(input CreateData) (*Data, error) {
	variables := map[string]any{
		"name":             input.Name,
		"keystoreId":       input.KeystoreId,
		"appCredentialsId": input.AppCredentialsId,
	}

	var response createResponse
	err := service.graphql.Query(createQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	data := response.AndroidAppBuildCredentials.Data
	return &Data{
		Id:               data.Id,
		Name:             data.Name,
		KeystoreId:       data.Keystore.Id,
		AppCredentialsId: input.AppCredentialsId,
	}, nil
}
