package googleserviceaccountkey

import (
	"fmt"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
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

	return findByProjectIdentifier(response.Account.ById.Data, getData.ProjectIdentifier, getData.ClientEmail)
}

func findByProjectIdentifier(keys []Data, projectIdentifier string, clientEmail string) (*Data, error) {
	var matches []Data
	for _, k := range keys {
		if k.ProjectIdentifier != projectIdentifier {
			continue
		}
		if clientEmail != "" && k.ClientEmail != clientEmail {
			continue
		}
		matches = append(matches, k)
	}

	switch len(matches) {
	case 0:
		if clientEmail != "" {
			return nil, fmt.Errorf("couldn't find google service account key with project identifier %s and client email %s", projectIdentifier, clientEmail)
		}
		return nil, fmt.Errorf("couldn't find google service account key with project identifier %s", projectIdentifier)
	case 1:
		return &matches[0], nil
	default:
		return nil, fmt.Errorf("found %d google service account keys with project identifier %s; set client_email to pick one", len(matches), projectIdentifier)
	}
}

type getByIdQuery struct {
	Data Data `json:"byId"`
}

type getByIdResponse struct {
	GoogleServiceAccountKey getByIdQuery `json:"googleServiceAccountKey"`
}

const getByIdQueryString = `
	query ($id: ID!) {
		googleServiceAccountKey {
			byId(id: $id) {
				id
				clientEmail
				clientIdentifier
				projectIdentifier
				privateKeyIdentifier
			}
		}
	}
`

func (s *service) Get(id string) (*Data, error) {
	variables := map[string]any{"id": id}

	var response getByIdResponse

	err := s.graphql.Query(getByIdQueryString, variables, &response)
	if err != nil {
		return nil, err
	}

	return &response.GoogleServiceAccountKey.Data, nil
}
