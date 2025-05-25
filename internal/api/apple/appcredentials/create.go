package appcredentials

type createIosAppCredentials struct {
	Data data `json:"createIosAppCredentials"`
}

type createResponse struct {
	IosAppCredentials createIosAppCredentials `json:"iosAppCredentials"`
}

const createQuery = `
	mutation ($appId: ID!, $appIdentifierId: ID!, $appStoreApiKeyId: ID $pushKeyId: ID) {
  iosAppCredentials {
    createIosAppCredentials(
      appId: $appId
      appleAppIdentifierId: $appIdentifierId
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

func (service *service) Create(input CreateData) (*Data, error) {
	variables := map[string]any{
		"appId":            input.AppId,
		"appIdentifierId":  input.AppIdentifierId,
		"appStoreApiKeyId": input.AppStoreApiKeyId,
		"pushKeyId":        input.PushKeyId,
	}

	var response createResponse

	err := service.graphql.Query(createQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	data := mapData(response.IosAppCredentials.Data)

	return &data, nil
}
