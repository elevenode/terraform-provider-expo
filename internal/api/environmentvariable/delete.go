package environmentvariable

const deleteQuery = `
	mutation ($id: ID!) {
		environmentVariable {
			deleteEnvironmentVariable(id: $id) {
				id
			}
		}
	}`

// Delete removes an environment variable from EAS
func (service *baseService) Delete(id string) (*any, error) {
	variables := map[string]any{"id": id}

	var response any

	err := service.graphql.Query(deleteQuery, variables, &response)
	return (*any)(nil), err
}
