package team

type getTeam struct {
	Data *data `json:"byAppleTeamIdentifier"`
}

type getTeamResponse struct {
	GetTeam getTeam `json:"appleTeam"`
}

const getQuery = `
	query ($identifier: String!, $accountId: ID!) {
		appleTeam {
			byAppleTeamIdentifier(accountId: $accountId, identifier: $identifier) {
				id
				appleTeamName
				appleTeamType
				appleTeamIdentifier
				account {
					id
				}
			}
		}
	}
`

func (service *service) GetByIdentifier(getData GetByIdentifierData) (*Data, error) {
	variables := map[string]any{
		"accountId":  getData.AccountId,
		"identifier": getData.Identifier,
	}

	var response getTeamResponse

	err := service.graphql.Query(getQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	return &Data{
		Id:         response.GetTeam.Data.Id,
		Name:       response.GetTeam.Data.Name,
		Identifier: response.GetTeam.Data.Identifier,
		Type:       response.GetTeam.Data.Type,
		AccountId:  response.GetTeam.Data.Account.Id,
	}, nil
}
