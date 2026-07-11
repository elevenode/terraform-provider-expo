package appcredentials

import "fmt"

type appByAppId struct {
	Data []data `json:"iosAppCredentials"`
}

type getResponse struct {
	AppByAppId appByAppId `json:"appByAppId"`
}

const getQuery = `
	query ($appId: String!) {
    appByAppId(appId: $appId) {
      iosAppCredentials {
        id
        appleAppIdentifier {
          id
        }
        app {
          id
        }
        appStoreConnectApiKeyForSubmissions {
          id
        }
        pushKey {
          id
        }
      	iosAppBuildCredentialsArray {
      		id
      		iosDistributionType
          provisioningProfile {
            id
          }
          distributionCertificate {
             id
          }
          iosAppCredentials {
            id
          }
      	}
      }
    }
	}
`

func (service *service) Get(data GetData) (*Data, error) {
	variables := map[string]any{
		"appId": data.AppId,
		"id":    data.Id,
	}

	var response getResponse

	err := service.graphql.Query(getQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	return findById(response.AppByAppId.Data, data.Id)
}

func findById(data []data, id string) (*Data, error) {
	for _, d := range data {
		if d.Id == id {
			data := mapData(d)
			return &data, nil
		}
	}
	return nil, fmt.Errorf("couldn't find app credentials with id %s", id)
}
