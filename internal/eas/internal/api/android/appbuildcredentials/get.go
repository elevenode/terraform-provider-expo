package appbuildcredentials

import "fmt"

type androidAppCredentials struct {
	Id   string `json:"id"`
	Data []data `json:"androidAppBuildCredentialsArray"`
}

type appByAppId struct {
	AndroidAppCredentials []androidAppCredentials `json:"androidAppCredentials"`
}

type getResponse struct {
	AppByAppId appByAppId `json:"appByAppId"`
}

const getQuery = `
	query ($appId: String!) {
		appByAppId(appId: $appId) {
			androidAppCredentials {
				id
				androidAppBuildCredentialsArray {
					name
					id
					androidKeystore { id }
				}
			}
		}
	}
`

func (service *service) Get(input GetData) (*Data, error) {
	variables := map[string]any{"appId": input.AppId}

	var response getResponse
	err := service.graphql.Query(getQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	appCredentials, err := findAndroidAppCredentialsById(response.AppByAppId.AndroidAppCredentials, input.AppCredentialsId)

	if err != nil {
		return nil, err
	}

	data, err := findById(appCredentials.Data, input.Id)
	if err != nil {
		return nil, err
	}
	return &Data{
		Id:               data.Id,
		Name:             data.Name,
		KeystoreId:       data.Keystore.Id,
		AppCredentialsId: appCredentials.Id,
	}, nil
}

func findAndroidAppCredentialsById(data []androidAppCredentials, id string) (*androidAppCredentials, error) {
	for _, item := range data {
		if item.Id == id {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("android app credentials with id %s not found", id)
}
func findById(data []data, id string) (*data, error) {
	for _, item := range data {
		if item.Id == id {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("android app build credentials with id %s not found", id)
}
