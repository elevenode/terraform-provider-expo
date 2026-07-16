package updatebranch

import "fmt"

type createUpdateBranch struct {
	Data *Data `json:"createUpdateBranchForApp"`
}

type createResponse struct {
	UpdateBranch createUpdateBranch `json:"updateBranch"`
}

const createQuery = `
	mutation ($appId: ID!, $name: String!) {
		updateBranch {
			createUpdateBranchForApp(appId: $appId, name: $name) {
				id
				name
			}
		}
	}`

// Creates an EAS Update Branch for an app
func (service *service) Create(data CreateData) (*Data, error) {
	variables := map[string]any{
		"appId": data.AppId,
		"name":  data.Name,
	}

	var response createResponse
	err := service.graphql.Query(createQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	if response.UpdateBranch.Data == nil {
		return nil, fmt.Errorf("couldn't create branch with name %s on app %s", data.Name, data.AppId)
	}

	return response.UpdateBranch.Data, nil
}
