package updatechannel

import "fmt"

type createUpdateChannel struct {
	Data *Data `json:"createUpdateChannelForApp"`
}

type createResponse struct {
	UpdateChannel createUpdateChannel `json:"updateChannel"`
}

const createQuery = `
	mutation ($appId: ID!, $name: String!, $branchMapping: String!) {
		updateChannel {
			createUpdateChannelForApp(appId: $appId, name: $name, branchMapping: $branchMapping) {
				id
				name
				branchMapping
			}
		}
	}`

// Creates an EAS Update Channel for an app
func (service *service) Create(data CreateData) (*Data, error) {
	variables := map[string]any{
		"appId":         data.AppId,
		"name":          data.Name,
		"branchMapping": data.BranchMapping,
	}

	var response createResponse
	err := service.graphql.Query(createQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	if response.UpdateChannel.Data == nil {
		return nil, fmt.Errorf("couldn't create channel with name %s on app %s", data.Name, data.AppId)
	}

	return response.UpdateChannel.Data, nil
}
