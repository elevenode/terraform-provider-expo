package appvariable

type updateEnvironmentVariable struct {
	Data *Data `json:"updateEnvironmentVariable"`
}

type updateResponse struct {
	UpdateAppVariable updateEnvironmentVariable `json:"environmentVariable"`
}

const updateQuery = `
	mutation ($id: ID!, $name: String!, $value: String!, $visibility: EnvironmentVariableVisibility!, $environments: [EnvironmentVariableEnvironment!]) {
		environmentVariable {
			updateEnvironmentVariable(
				environmentVariableData: { environments: $environments, id: $id, name: $name, value: $value, visibility: $visibility }
			) {
				environments
				id
				name
				value
				visibility
			}
		}
	}`

// Updates an App Environment Variable in EAS
func (service *service) Update(data UpdateData) (*Data, error) {
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

	return response.UpdateAppVariable.Data, err
}
