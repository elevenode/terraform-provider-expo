package team

type updateTeam struct {
	Data *data `json:"updateAppleTeam"`
}

type updateTeamResponse struct {
	UpdateAppleTeam updateTeam `json:"appleTeam"`
}

const updateQuery = `
	mutation ($id: ID!, $name: String!, $type: AppleTeamType!) {
		appleTeam {
			updateAppleTeam(
				appleTeamUpdateInput: { appleTeamName: $name, appleTeamType: $type }
				id: $id
			) {
				appleTeamName
				appleTeamType
				id
				appleTeamIdentifier
				account {
					id
				}
			}
		}
	}`

func (service *service) Update(data UpdateData) (*Data, error) {
	variables := map[string]any{
		"id":   data.Id,
		"name": data.Name,
		"type": data.Type,
	}

	var response updateTeamResponse

	err := service.graphql.Query(updateQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	return &Data{
		Id:         response.UpdateAppleTeam.Data.Id,
		Name:       response.UpdateAppleTeam.Data.Name,
		Identifier: response.UpdateAppleTeam.Data.Identifier,
		Type:       response.UpdateAppleTeam.Data.Type,
		AccountId:  response.UpdateAppleTeam.Data.Account.Id,
	}, nil
}
