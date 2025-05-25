package appcredentials

type updateIosAppCredentials struct {
	Data data `json:"updateIosAppCredentials"`
}

type updateResponse struct {
	IosAppCredentials updateIosAppCredentials `json:"iosAppCredentials"`
}

const updateQuery = `
	mutation ($id: ID!, $appStoreApiKeyId: ID, $pushKeyId: ID) {
    iosAppCredentials {
      updateIosAppCredentials(
        id: $id
        iosAppCredentialsInput: { appStoreConnectApiKeyForSubmissionsId: $appStoreApiKeyId, pushKeyId: $pushKeyId }
      ) {
        id
        app { id }
        appleAppIdentifier { id }
        appStoreConnectApiKeyForSubmissions { id }
        pushKey { id }
        iosAppBuildCredentialsArray {
          id
          iosDistributionType
          provisioningProfile { id }
          distributionCertificate { id }
          iosAppCredentials { id }
        }
      }
    }
	}
`

func (service *service) Update(input UpdateData) (*Data, error) {
	variables := map[string]any{
		"id":               input.Id,
		"appStoreApiKeyId": input.AppStoreApiKeyId,
		"pushKeyId":        input.PushKeyId,
	}

	var response updateResponse

	err := service.graphql.Query(updateQuery, variables, &response)
	if err != nil {
		return nil, err
	}

	data := mapData(response.IosAppCredentials.Data)

	return &data, nil
}
