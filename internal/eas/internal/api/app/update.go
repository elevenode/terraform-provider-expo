package app

type updateApp struct {
	Data *Data `json:"setAppInfo"`
}

type updateResponse struct {
	UpdateApp updateApp `json:"app"`
}

const updateQuery = `
	mutation ($id: ID!, $name: String!) {
		app {
			setAppInfo(appId: $id, appInfo: { displayName: $name }) {
				name
				slug
				id
			}
		}
	}`

// Creates an App in EAS
func (service *service) Update(data UpdateData) (*Data, error) {
	variables := map[string]any{
		"id":   data.Id,
		"name": data.Name,
	}

	var response updateResponse
	err := service.graphql.Query(updateQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	return response.UpdateApp.Data, nil
}
