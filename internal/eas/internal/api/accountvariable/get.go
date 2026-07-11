package accountvariable

import "fmt"

type accountById struct {
	Data []Data `json:"environmentVariablesIncludingSensitive"`
}

type accountQuery struct {
	ById accountById `json:"byId"`
}

type getResponse struct {
	Account accountQuery `json:"account"`
}

const getQuery = `
	query ($accountId: String!) {
		account {
			byId(accountId: $accountId) {
				environmentVariablesIncludingSensitive {
					createdAt
					environments
					id
					name
					updatedAt
					value
					visibility
				}
			}
		}
	}`

func (service *service) getAccountEnvironmentVariables(accountId string) ([]Data, error) {
	variables := map[string]any{"accountId": accountId}

	var response getResponse

	err := service.graphql.Query(getQuery, variables, &response)
	return response.Account.ById.Data, err
}

func (service *service) Get(getData GetData) (*Data, error) {
	data, err := service.getAccountEnvironmentVariables(getData.AccountId)
	if err != nil {
		return nil, err
	}

	return findById(data, getData.Id)
}

func findById(variables []Data, id string) (*Data, error) {
	for _, variable := range variables {
		if variable.Id == id {
			return &variable, nil
		}
	}
	return nil, fmt.Errorf("couldn't find variable with id %s", id)
}
