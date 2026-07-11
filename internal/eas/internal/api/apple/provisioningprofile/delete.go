package provisioningprofile

const deleteQuery = `
	mutation ($id: ID!) {
    appleProvisioningProfile {
      deleteAppleProvisioningProfile(id: $id) {
        id
      }
    }
	}`

func (service *service) Delete(id string) (*any, error) {
	variables := map[string]any{"id": id}

	var response any

	err := service.graphql.Query(deleteQuery, variables, &response)
	return (*any)(nil), err
}
