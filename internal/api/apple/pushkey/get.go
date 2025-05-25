package pushkey

import (
	"fmt"

	"github.com/fintreal/eas-sdk-go/internal/utils"
)

type getAccount struct {
	Data []Data `json:"applePushKeys"`
}

type getResponse = utils.AccountResponse[getAccount]

const getQuery = `
	query ($accountId: String!) {
    account {
      byId(accountId: $accountId) {
        applePushKeys {
          id
          keyIdentifier
        }
      }
    }
	}
`

func (s *service) GetByIdentifier(getData GeyByIdentifierData) (*Data, error) {
	variables := map[string]any{"accountId": getData.AccountId}

	var response getResponse

	err := s.graphql.Query(getQuery, variables, &response)
	if err != nil {
		return nil, err
	}

	return findByIdentifier(response.Account.ById.Data, getData.Identifier)
}

func findByIdentifier(apiKeys []Data, identifier string) (*Data, error) {
	for _, apiKey := range apiKeys {
		if apiKey.Identifier == identifier {
			return &apiKey, nil
		}
	}
	return nil, fmt.Errorf("couldn't find push key with identifier %s", identifier)
}
