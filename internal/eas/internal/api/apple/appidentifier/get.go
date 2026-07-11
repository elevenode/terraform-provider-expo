package appidentifier

import (
	"fmt"

	"github.com/fintreal/eas-sdk-go/internal/utils"
)

type getResponse struct {
	Data []Data `json:"appleAppIdentifiers"`
}

const getQuery = `
	query ($accountId: String!, $identifier: String!) {
		account {
			byId(accountId: $accountId) {
				appleAppIdentifiers(bundleIdentifier: $identifier) {
					bundleIdentifier
					id
				}
			}
		}
	}`

func (s *service) GetByIdentifier(getData GetByIdentifierData) (*Data, error) {
	variables := map[string]any{
		"identifier": getData.Identifier,
		"accountId":  getData.AccountId,
	}

	var response utils.AccountResponse[getResponse]

	err := s.graphql.Query(getQuery, variables, &response)
	if err != nil {
		return nil, err
	}
	return findBundleIdentifierByIdentifier(response.Account.ById.Data, getData.Identifier)
}

func findBundleIdentifierByIdentifier(identifiers []Data, identifier string) (*Data, error) {
	for _, bundleIdentifier := range identifiers {
		if bundleIdentifier.Identifier == identifier {
			return &bundleIdentifier, nil
		}
	}
	return nil, fmt.Errorf("couldn't find bundle identifier with identifier %s", identifier)
}
