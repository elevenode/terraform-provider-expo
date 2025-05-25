package app

import "fmt"

type getResponse struct {
	Data *Data `json:"appByAppId"`
}

const getQuery = `
	query App ($id: String!) {
    	appByAppId(appId: $id) {
        	id
       		name
        	slug
    	}
	}`

// Retrieves an App from EAS by it's id
func (service *service) Get(id string) (*Data, error) {
	variables := map[string]any{"id": id}

	var response getResponse
	err := service.graphql.Query(getQuery, variables, &response)

	return response.Data, err
}

const getByFullNameQuery = `
	query App($fullName: String!) {
    app {
      byFullName(fullName: $fullName) {
        id
        name
        slug
      }
    }
	}
`

type app struct {
	Data *Data `json:"byFullName"`
}

type getByFullNameResponse struct {
	App *app `json:"app"`
}

// Retrieves an App from EAS by it's full name (@organization/app-slug)
func (service *service) GetByFullName(fullName string) (*Data, error) {
	variables := map[string]any{"fullName": fullName}

	var response getByFullNameResponse
	err := service.graphql.Query(getByFullNameQuery, variables, &response)

	if response.App == nil {
		return nil, fmt.Errorf("app not found")
	}

	return response.App.Data, err
}
