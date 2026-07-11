package team

type createTeam struct {
	Data *data `json:"createAppleTeam"`
}

type createTeamResponse struct {
	CreateTeam createTeam `json:"appleTeam"`
}

const createQuery = `
	mutation ($accountId: ID!, $name: String!, $identifier: String!, $type: AppleTeamType!) {
		appleTeam {
			createAppleTeam(
				accountId: $accountId
				appleTeamInput: {
					appleTeamName: $name
					appleTeamType: $type
					appleTeamIdentifier: $identifier
				}
			) {
				appleTeamIdentifier
				appleTeamName
				appleTeamType
				id
				account {
					id
				}
			}
		}
	}`

func (service *service) Create(data CreateData) (*Data, error) {
	variables := map[string]any{
		"identifier": data.Identifier,
		"name":       data.Name,
		"accountId":  data.AccountId,
		"type":       data.Type,
	}

	var response createTeamResponse

	err := service.graphql.Query(createQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	return &Data{
		Id:         response.CreateTeam.Data.Id,
		Name:       response.CreateTeam.Data.Name,
		Identifier: response.CreateTeam.Data.Identifier,
		Type:       response.CreateTeam.Data.Type,
		AccountId:  response.CreateTeam.Data.Account.Id,
	}, nil
}
