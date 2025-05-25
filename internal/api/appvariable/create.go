package appvariable

type createEnvironmentVariable struct {
	Data *Data `json:"createEnvironmentVariableForApp"`
}

type createResponse struct {
	EnvironmentVariable createEnvironmentVariable `json:"environmentVariable"`
}

const createQuery = `
	mutation ($appId: ID!, $name: String!, $value: String!, $visibility: EnvironmentVariableVisibility!, $environments: [EnvironmentVariableEnvironment!]) {
		environmentVariable {
			createEnvironmentVariableForApp(
				appId: $appId
				environmentVariableData: {
					name: $name
					value: $value
					type: STRING
					visibility: $visibility
					environments: $environments
				}
			) {
				environments
				id
				name
				scope
				type
				value
				visibility
			}
		}
	}`

// Creates an App Environment Variable in EAS
func (service *service) Create(data CreateData) (*Data, error) {
	variables := map[string]any{
		"appId":        data.AppId,
		"name":         data.Name,
		"value":        data.Value,
		"visibility":   data.Visibility,
		"environments": data.Environments,
	}

	var response createResponse
	err := service.graphql.Query(createQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	return response.EnvironmentVariable.Data, err
}
