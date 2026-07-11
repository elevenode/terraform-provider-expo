package appcredentials

type createAndroidAppCredentials struct {
	Data data `json:"createAndroidAppCredentials"`
}

type createResponse struct {
	AndroidAppCredentials createAndroidAppCredentials `json:"androidAppCredentials"`
}

const createQuery = `
	mutation ($appId: ID!, $identifier: String!, $googleServiceAccountKeyId: ID!) {
    androidAppCredentials {
      createAndroidAppCredentials(
        appId: $appId
        applicationIdentifier: $identifier
        androidAppCredentialsInput: { googleServiceAccountKeyForSubmissionsId: $googleServiceAccountKeyId }
      ) {
        id
        applicationIdentifier
        googleServiceAccountKeyForSubmissions { id }
        app { id }
				androidAppBuildCredentialsArray {
					name
					id
					androidKeystore { id }
				}
      }
    }
	}
`

func (service *service) Create(input CreateData) (*Data, error) {
	variables := map[string]any{
		"appId":                     input.AppId,
		"identifier":                input.Identifier,
		"googleServiceAccountKeyId": input.GoogleServiceAccountKeyId,
	}

	var response createResponse

	err := service.graphql.Query(createQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	data := mapData(response.AndroidAppCredentials.Data)

	return &data, nil
}
