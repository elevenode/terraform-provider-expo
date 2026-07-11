package appbuildcredentials

type createIosAppBuildCredentials struct {
	Data data `json:"createIosAppBuildCredentials"`
}

type createResponse struct {
	IosAppBuildCredentials createIosAppBuildCredentials `json:"iosAppBuildCredentials"`
}

const createQuery = `
	mutation ($certificateId: ID!, $distributionType: IosDistributionType!, $provisioningProfileId: ID!, $appCredentialsId: ID!) {
    iosAppBuildCredentials {
      createIosAppBuildCredentials(
        iosAppBuildCredentialsInput: {
            distributionCertificateId: $certificateId
            iosDistributionType: $distributionType
            provisioningProfileId: $provisioningProfileId
        }
        iosAppCredentialsId: $appCredentialsId
      ) {
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
`

func (service *service) Create(input CreateData) (*Data, error) {
	variables := map[string]any{
		"distributionType":      input.DistributionType,
		"certificateId":         input.CertificateId,
		"provisioningProfileId": input.ProvisioningProfileId,
		"appCredentialsId":      input.AppCredentialsId,
	}

	var response createResponse

	err := service.graphql.Query(createQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	data := response.IosAppBuildCredentials.Data

	return &Data{
		Id:                    data.Id,
		DistributionType:      data.DistributionType,
		CertificateId:         data.Certificate.Id,
		ProvisioningProfileId: data.ProvisioningProfile.Id,
		AppCredentialsId:      data.AppCredentials.Id,
	}, nil
}
