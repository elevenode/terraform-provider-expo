package certificate

import "github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"

type getResponse struct {
	Data []Data `json:"appleDistributionCertificates"`
}

const getQuery = `
	query ($accountId: String!) {
			account {
					byId(accountId: $accountId) {
							appleDistributionCertificates {
									id
									serialNumber
							}
					}
			}
	}`

func (service *service) GetBySerialNumber(getData GetBySerialNumberData) (*Data, error) {
	variables := map[string]any{
		"accountId":    getData.AccountId,
		"serialNumber": getData.SerialNumber,
	}

	var response utils.AccountResponse[getResponse]

	err := service.graphql.Query(getQuery, variables, &response)
	if err != nil {
		return nil, err
	}
	return findBySerialNumber(response.Account.ById.Data, getData.SerialNumber)
}

func findBySerialNumber(certificates []Data, serialNumber string) (*Data, error) {
	for _, certificate := range certificates {
		if certificate.SerialNumber == serialNumber {
			return &certificate, nil
		}
	}
	return nil, nil
}
