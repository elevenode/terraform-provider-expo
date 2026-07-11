package account

type account struct {
	ByName *Data `json:"byName"`
}

type getResponse struct {
	Account account `json:"account"`
}

const getByNameQuery = `
	query ($name: String!) {
		account {
			byName(accountName: $name) {
				id
				name
			}
		}
	}`

// Retrieves the currents user's organization account by it's name
//
// @name organization name
func (service *accountService) GetByName(name string) (*Data, error) {
	variables := map[string]any{"name": name}

	var response getResponse

	err := service.graphql.Query(getByNameQuery, variables, &response)

	return response.Account.ByName, err
}
