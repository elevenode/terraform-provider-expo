package accountvariable

type createEnvironmentVariable struct {
	Data *Data `json:"createEnvironmentVariableForAccount"`
}

type createResponse struct {
	EnvironmentVariable createEnvironmentVariable `json:"environmentVariable"`
}

const createQuery = `
	mutation ($accountId: ID!, $name: String!, $value: String!, $visibility: EnvironmentVariableVisibility!, $environments: [EnvironmentVariableEnvironment!]) {
		environmentVariable {
			createEnvironmentVariableForAccount(
				accountId: $accountId
				environmentVariableData: {
					name: $name
					value: $value
					type: STRING
					visibility: $visibility
					environments: $environments
				}
			) {
				createdAt
				environments
				id
				name
				scope
				type
				updatedAt
				value
				visibility
			}
		}
	}`

func (service *service) Create(data CreateData) (*Data, error) {
	variables := map[string]any{
		"accountId":    data.AccountId,
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
