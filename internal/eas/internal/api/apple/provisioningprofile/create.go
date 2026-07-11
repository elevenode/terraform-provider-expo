package provisioningprofile

type appleProvisioningProfile struct {
	Data data `json:"createAppleProvisioningProfile"`
}

type createResponse struct {
	AppleProvisioningProfile appleProvisioningProfile `json:"appleProvisioningProfile"`
}

const createQuery = `
    mutation ($accountId: ID!, $appleAppIdentifierId: ID!, $base64: String!) {
        appleProvisioningProfile {
            createAppleProvisioningProfile(
                accountId: $accountId
                appleAppIdentifierId: $appleAppIdentifierId
                appleProvisioningProfileInput: { appleProvisioningProfile: $base64 }
            ) {
                id
                provisioningProfile
                appleAppIdentifier {
                    id
                }
            }
        }
    }
`

func (service *service) Create(data CreateData) (*Data, error) {
	variables := map[string]any{
		"accountId":            data.AccountId,
		"appleAppIdentifierId": data.AppBundleIdentifierId,
		"base64":               data.Base64,
	}

	var response createResponse

	err := service.graphql.Query(createQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	return &Data{
		Id:                    response.AppleProvisioningProfile.Data.Id,
		AppBundleIdentifierId: response.AppleProvisioningProfile.Data.AppleAppIdentifier.Id,
		Base64:                response.AppleProvisioningProfile.Data.Base64,
	}, nil
}
