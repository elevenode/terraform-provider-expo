package googleserviceaccountkey

const deleteQuery = `
	mutation ($id: ID!) {
		googleServiceAccountKey {
			deleteGoogleServiceAccountKey(id: $id) {
				id
			}
		}
	}
`

func (s *service) Delete(id string) (*any, error) {
	variables := map[string]any{"id": id}

	var response any
	err := s.graphql.Query(deleteQuery, variables, &response)

	return nil, err
}
