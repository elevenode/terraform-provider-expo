package app

import "github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/backgroundjob"

type scheduleAppDeletion struct {
	Data *backgroundjob.Data `json:"scheduleAppDeletion"`
}

type deleteResponse struct {
	App scheduleAppDeletion `json:"app"`
}

const deleteQuery = `
	mutation ($id: ID!) {
		app {
			scheduleAppDeletion(appId: $id) {
				id
				state
				tries
				willRetry
				resultId
				resultType
				errorCode
				errorMessage
				createdAt
				updatedAt
			}
		}
	}`

// Schedules an App for deletion in EAS. Deletion runs asynchronously; poll the returned
// background job receipt to observe whether it succeeded.
func (service *service) Delete(id string) (*backgroundjob.Data, error) {
	variables := map[string]any{"id": id}

	var response deleteResponse
	err := service.graphql.Query(deleteQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	return response.App.Data, nil
}
