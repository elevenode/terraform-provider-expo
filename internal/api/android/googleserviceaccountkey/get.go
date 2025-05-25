package googleserviceaccountkey

import (
	"fmt"

	"github.com/fintreal/eas-sdk-go/internal/utils"
)

type getAccount struct {
	Data []Data `json:"googleServiceAccountKeys"`
}

type getResponse = utils.AccountResponse[getAccount]

const getQuery = `
	query ($accountId: String!) {
    account {
      byId(accountId: $accountId) {
        googleServiceAccountKeys {
          id
          clientEmail
          clientIdentifier
          projectIdentifier
        }
      }
    }
	}
`

func (s *service) GetByProjectIdentifier(getData GetByProjectIdentifierData) (*Data, error) {
	variables := map[string]any{"accountId": getData.AccountId}

	var response getResponse

	err := s.graphql.Query(getQuery, variables, &response)
	if err != nil {
		return nil, err
	}

	return findByProjectIdentifier(response.Account.ById.Data, getData.ProjectIdentifier)
}
func findByProjectIdentifier(identifiers []Data, identifier string) (*Data, error) {
	for _, i := range identifiers {
		if i.ProjectIdentifier == identifier {
			return &i, nil
		}
	}
	return nil, fmt.Errorf("couldn't find google service account key with project identifier %s", identifier)
}
