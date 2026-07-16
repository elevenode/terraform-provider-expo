package updatechannel

import "fmt"

type editUpdateChannel struct {
	Data *Data `json:"editUpdateChannel"`
}

type updateResponse struct {
	UpdateChannel editUpdateChannel `json:"updateChannel"`
}

const updateQuery = `
	mutation ($channelId: ID!, $branchMapping: String!) {
		updateChannel {
			editUpdateChannel(channelId: $channelId, branchMapping: $branchMapping) {
				id
				name
				branchMapping
			}
		}
	}`

// Updates the branch mapping of an EAS Update Channel
func (service *service) Update(data UpdateData) (*Data, error) {
	variables := map[string]any{
		"channelId":     data.Id,
		"branchMapping": data.BranchMapping,
	}

	var response updateResponse
	err := service.graphql.Query(updateQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	if response.UpdateChannel.Data == nil {
		return nil, fmt.Errorf("couldn't find channel with id %s", data.Id)
	}

	return response.UpdateChannel.Data, nil
}
