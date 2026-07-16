package updatebranch

const deleteQuery = `
	mutation ($branchId: ID!) {
		updateBranch {
			deleteUpdateBranch(branchId: $branchId) {
				id
			}
		}
	}`

// Delete removes an EAS Update Branch from EAS
func (service *service) Delete(id string) (*any, error) {
	variables := map[string]any{"branchId": id}

	var response any

	err := service.graphql.Query(deleteQuery, variables, &response)
	return (*any)(nil), err
}
