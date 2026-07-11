package appbuildcredentials

import "fmt"

type iosAppCredentials struct {
	Id   string `json:"id"`
	Data []data `json:"iosAppBuildCredentialsList"`
}

type appByAppId struct {
	IosAppCredentials []iosAppCredentials `json:"iosAppCredentials"`
}

type getResponse struct {
	AppByAppId appByAppId `json:"appByAppId"`
}

const getQuery = `
	query ($appId: String!) {
    appByAppId(appId: $appId) {
      iosAppCredentials {
				id
        iosAppBuildCredentialsList {
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
	variables := map[string]any{"appId": data.AppId}

	var response getResponse
	err := service.graphql.Query(getQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	appCredentials, err := findIosAppCredentialsById(response.AppByAppId.IosAppCredentials, data.AppCredentialsId)

	if err != nil {
		return nil, err
	}

	return findById(appCredentials.Data, data.Id)
}

func findIosAppCredentialsById(data []iosAppCredentials, id string) (*iosAppCredentials, error) {
	for _, d := range data {
		if d.Id == id {
			return &d, nil
		}
	}
	return nil, fmt.Errorf("couldn't find app build credentials with id %s", id)
}

func findById(data []data, id string) (*Data, error) {
	for _, d := range data {
		if d.Id == id {
			return &Data{
				Id:                    d.Id,
				DistributionType:      d.DistributionType,
				CertificateId:         d.Certificate.Id,
				ProvisioningProfileId: d.ProvisioningProfile.Id,
				AppCredentialsId:      d.AppCredentials.Id,
			}, nil
		}
	}
	return nil, fmt.Errorf("couldn't find app build credentials with id %s", id)
}
