package appvariable

const deleteQuery = `
	mutation ($id: ID!) {
		environmentVariable {
			deleteEnvironmentVariable(id: $id) {
				id
			}
		}
	}`

// Deletes an App Environment Variable from EAS
func (service *service) Delete(id string) (*any, error) {
	variables := map[string]any{"id": id}

	var response any

	err := service.graphql.Query(deleteQuery, variables, &response)
	return (*any)(nil), err
}
