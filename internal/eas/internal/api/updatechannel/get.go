package updatechannel

import "fmt"

type channelByAppId struct {
	Data *Data `json:"updateChannelByName"`
}

type appQuery struct {
	ById channelByAppId `json:"byId"`
}

type getResponse struct {
	App appQuery `json:"app"`
}

const getByNameQuery = `
	query ($appId: String!, $channelName: String!) {
		app {
			byId(appId: $appId) {
				id
				updateChannelByName(name: $channelName) {
					id
					name
					branchMapping
				}
			}
		}
	}`

// Retrieves an EAS Update Channel by its name and appId
func (service *service) GetByName(data GetByNameData) (*Data, error) {
	variables := map[string]any{
		"appId":       data.AppId,
		"channelName": data.Name,
	}

	var response getResponse
	err := service.graphql.Query(getByNameQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	if response.App.ById.Data == nil {
		return nil, fmt.Errorf("couldn't find channel with name %s on app %s", data.Name, data.AppId)
	}

	return response.App.ById.Data, nil
}
