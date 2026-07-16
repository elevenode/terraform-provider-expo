package updatebranch

import "fmt"

type branchByAppId struct {
	Data *Data `json:"updateBranchByName"`
}

type appQuery struct {
	ById branchByAppId `json:"byId"`
}

type getResponse struct {
	App appQuery `json:"app"`
}

const getByNameQuery = `
	query ($appId: String!, $name: String!) {
		app {
			byId(appId: $appId) {
				id
				updateBranchByName(name: $name) {
					id
					name
				}
			}
		}
	}`

// Retrieves an EAS Update Branch by its name and appId
func (service *service) GetByName(data GetByNameData) (*Data, error) {
	variables := map[string]any{
		"appId": data.AppId,
		"name":  data.Name,
	}

	var response getResponse
	err := service.graphql.Query(getByNameQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	if response.App.ById.Data == nil {
		return nil, fmt.Errorf("couldn't find branch with name %s on app %s", data.Name, data.AppId)
	}

	return response.App.ById.Data, nil
}
