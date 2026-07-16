package backgroundjob

import "fmt"

type receiptById struct {
	Data *Data `json:"byId"`
}

type getResponse struct {
	BackgroundJobReceipt receiptById `json:"backgroundJobReceipt"`
}

const getByIdQuery = `
	query ($id: ID!) {
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
			}
		}
	}`

// Retrieves a background job receipt by its id
func (service *service) GetById(id string) (*Data, error) {
	variables := map[string]any{"id": id}

	var response getResponse
	err := service.graphql.Query(getByIdQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	if response.BackgroundJobReceipt.Data == nil {
		return nil, fmt.Errorf("couldn't find background job receipt with id %s", id)
	}

	return response.BackgroundJobReceipt.Data, nil
}
