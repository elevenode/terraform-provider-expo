package updatechannel

const deleteQuery = `
	mutation ($channelId: ID!) {
		updateChannel {
			deleteUpdateChannel(channelId: $channelId) {
				id
			}
		}
	}`

// Delete removes an EAS Update Channel from EAS. The channel must not point at
// any branches; clear its branch mapping first.
func (service *service) Delete(id string) (*any, error) {
	variables := map[string]any{"channelId": id}

	var response any

	err := service.graphql.Query(deleteQuery, variables, &response)
	return (*any)(nil), err
}
