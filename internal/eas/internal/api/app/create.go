package app

type createApp struct {
	Data *Data `json:"createApp"`
}

type createResponse struct {
	CreateApp createApp `json:"app"`
}

const createQuery = `
	mutation ($accountId: ID!, $name: String!, $slug: String!) {
		app {
			createApp(
				appInput: {
					accountId: $accountId,
					appInfo: { displayName: $name }
					projectName: $slug
				}
			) {
				id
				name
				slug
			}
		}
	}`

// Creates an App in EAS
func (service *service) Create(data CreateData) (*Data, error) {
	variables := map[string]any{
		"accountId": data.AccountId,
		"name":      data.Name,
		"slug":      data.Slug,
	}

	var response createResponse
	err := service.graphql.Query(createQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	return response.CreateApp.Data, nil
}
