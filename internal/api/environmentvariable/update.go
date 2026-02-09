package environmentvariable

type updateEnvironmentVariable struct {
	Data *Data `json:"updateEnvironmentVariable"`
}

type updateResponse struct {
	EnvironmentVariable updateEnvironmentVariable `json:"environmentVariable"`
}

const updateQuery = `
	mutation ($id: ID!, $name: String!, $value: String!, $visibility: EnvironmentVariableVisibility!, $environments: [EnvironmentVariableEnvironment!]) {
		environmentVariable {
			updateEnvironmentVariable(
				environmentVariableData: { environments: $environments, id: $id, name: $name, value: $value, visibility: $visibility }
			) {
				createdAt
				environments
				id
				name
				updatedAt
				value
				visibility
			}
		}
	}`

func (service *baseService) Update(data UpdateData) (*Data, error) {
	variables := map[string]any{
		"id":           data.Id,
		"name":         data.Name,
		"value":        data.Value,
		"visibility":   data.Visibility,
		"environments": data.Environments,
	}

	var response updateResponse
	err := service.graphql.Query(updateQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	return response.EnvironmentVariable.Data, err
}
