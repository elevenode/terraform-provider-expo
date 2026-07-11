package me

import "github.com/fintreal/eas-sdk-go/internal/utils"

type getResponse struct {
	Data *Data `json:"meActor"`
}

const getQuery = `
	query {
		meActor {
			displayName
			id
		}
	}`

// Retrieves the current user name and id from EAS
func (service *service) Get() (*Data, error) {
	inputMapper := func(a any) map[string]any {
		return map[string]any{}
	}

	outputMapper := func(response getResponse, _ any) *Data {
		return response.Data
	}

	config := utils.Configuration[any, Data, getResponse]{
		Query:        getQuery,
		InputMapper:  inputMapper,
		OutputMapper: outputMapper,
	}

	return utils.ConfigurationFunc(config, service.graphql)(nil)
}
