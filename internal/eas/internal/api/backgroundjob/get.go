package backgroundjob

import "fmt"

type backgroundJobReceipt struct {
	Data *Data `json:"byId"`
}

type getResponse struct {
	BackgroundJobReceipt *backgroundJobReceipt `json:"backgroundJobReceipt"`
}

const getQuery = `
	query BackgroundJobReceipt ($id: ID!) {
		backgroundJobReceipt {
			byId(id: $id) {
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

// Retrieves a background job receipt from EAS by it's id
func (service *service) Get(id string) (*Data, error) {
	variables := map[string]any{"id": id}

	var response getResponse
	err := service.graphql.Query(getQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	if response.BackgroundJobReceipt == nil {
		return nil, fmt.Errorf("background job receipt not found")
	}

	return response.BackgroundJobReceipt.Data, nil
}
