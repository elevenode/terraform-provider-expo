package appidentifier

type createAppBundleIdentifier struct {
	Data Data `json:"createAppleAppIdentifier"`
}

type createResponse struct {
	CreateAppIdentifier createAppBundleIdentifier `json:"appleAppIdentifier"`
}

const createQuery = `
	mutation ($accountId: ID!, $identifier: String!) {
		appleAppIdentifier {
			createAppleAppIdentifier(
				accountId: $accountId
				appleAppIdentifierInput: { bundleIdentifier: $identifier }
			) {
				bundleIdentifier
				id
			}
		}
	}
`

func (s *service) Create(data CreateData) (*Data, error) {
	variables := map[string]any{
		"accountId":  data.AccountId,
		"identifier": data.Identifier,
	}

	var response createResponse
	err := s.graphql.Query(createQuery, variables, &response)
	if err != nil {
		return nil, err
	}

	return &response.CreateAppIdentifier.Data, nil
}
