package googleserviceaccountkey

import (
	"encoding/json"
	"fmt"
)

type createGoogleServiceAccountKey struct {
	Data Data `json:"createGoogleServiceAccountKey"`
}

type createResponse struct {
	GoogleServiceAccountKey createGoogleServiceAccountKey `json:"googleServiceAccountKey"`
}

const createQuery = `
	mutation ($accountId: ID!, $jsonKey: JSONObject!) {
		googleServiceAccountKey {
			createGoogleServiceAccountKey(
				accountId: $accountId
				googleServiceAccountKeyInput: { jsonKey: $jsonKey }
			) {
				id
				clientEmail
				clientIdentifier
				projectIdentifier
				privateKeyIdentifier
			}
		}
	}
`

// Uploads a Google service account key (the JSON key file) to the EAS account.
func (s *service) Create(input CreateData) (*Data, error) {
	// jsonKey is a JSONObject scalar, so it must be sent as an object, not a string.
	var jsonKey map[string]any
	if err := json.Unmarshal([]byte(input.KeyJson), &jsonKey); err != nil {
		return nil, fmt.Errorf("google service account key is not valid JSON: %w", err)
	}

	variables := map[string]any{
		"accountId": input.AccountId,
		"jsonKey":   jsonKey,
	}

	var response createResponse

	err := s.graphql.Query(createQuery, variables, &response)
	if err != nil {
		return nil, err
	}

	return &response.GoogleServiceAccountKey.Data, nil
}
