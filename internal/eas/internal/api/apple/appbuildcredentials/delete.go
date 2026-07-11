package appbuildcredentials

const deleteQuery = `
	mutation ($id: ID!) {
    iosAppBuildCredentials {
      deleteIosAppBuildCredentials(id: $id) {
        id
      }
    }
	}
`

func (service *service) Delete(id string) (*any, error) {
	variables := map[string]any{
		"id": id,
	}

	var response any

	err := service.graphql.Query(deleteQuery, variables, &response)

	return nil, err
}
