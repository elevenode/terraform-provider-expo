package provisioningprofile

import (
	"fmt"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

type getProvisioningProfilesResponse struct {
	Data []data `json:"appleProvisioningProfiles"`
}

const getQuery = `
	query ($accountId: String!) {
			account {
					byId(accountId: $accountId) {
							appleProvisioningProfiles {
									id
									provisioningProfile
									appleAppIdentifier {
											id
									}
							}
					}
			}
	}
`

func (service *service) Get(getData GetData) (*Data, error) {
	variables := map[string]any{
		"accountId": getData.AccountId,
	}

	var response utils.AccountResponse[getProvisioningProfilesResponse]
	err := service.graphql.Query(getQuery, variables, &response)
	if err != nil {
		return nil, err
	}
	return findProvisioningProfileById(response.Account.ById.Data, getData.Id)
}

func findProvisioningProfileById(provisioningProfiles []data, id string) (*Data, error) {
	for _, provisioningProfile := range provisioningProfiles {
		if provisioningProfile.Id == id {
			return &Data{
				Id:                    provisioningProfile.Id,
				AppBundleIdentifierId: provisioningProfile.AppleAppIdentifier.Id,
				Base64:                provisioningProfile.Base64,
			}, nil
		}
	}
	return nil, fmt.Errorf("couldn't find provisioning profile with id %s", id)
}
