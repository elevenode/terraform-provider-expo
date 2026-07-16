package updatechannel

import (
	"fmt"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/backgroundjob"
)

type scheduleUpdateChannelDeletion struct {
	Data *backgroundjob.Data `json:"scheduleUpdateChannelDeletion"`
}

type deleteResponse struct {
	UpdateChannel scheduleUpdateChannelDeletion `json:"updateChannel"`
}

const deleteQuery = `
	mutation ($channelId: ID!) {
		updateChannel {
			scheduleUpdateChannelDeletion(channelId: $channelId) {
				id
				state
				tries
				willRetry
				resultId
				resultType
				errorCode
				errorMessage
			}
		}
	}`

// Delete schedules deletion of an EAS Update Channel and returns the receipt of
// the background job, which must be polled for completion.
//
// The deprecated deleteUpdateChannel is unusable: it requires a channel that
// points at no branches, and EAS rejects editing a mapping to be empty. This
// also permanently deletes every build associated with the channel.
func (service *service) Delete(id string) (*backgroundjob.Data, error) {
	variables := map[string]any{"channelId": id}

	var response deleteResponse
	err := service.graphql.Query(deleteQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	if response.UpdateChannel.Data == nil {
		return nil, fmt.Errorf("couldn't schedule deletion of channel with id %s", id)
	}

	return response.UpdateChannel.Data, nil
}
