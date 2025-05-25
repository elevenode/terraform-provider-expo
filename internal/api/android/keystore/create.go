package keystore

type createAndroidKeystore struct {
	Data Data `json:"createAndroidKeystore"`
}

type createResponse struct {
	AndroidKeystore createAndroidKeystore `json:"androidKeystore"`
}

const createQuery = `
	mutation (
		$accountId: ID!
		$keystoreBase64: String!
		$keyAlias: String!
		$keyPassword: String!
		$keystorePassword: String!
	) {
		androidKeystore {
			createAndroidKeystore(
				accountId: $accountId
				androidKeystoreInput: {
					base64EncodedKeystore: $keystoreBase64
					keyAlias: $keyAlias
					keyPassword: $keyPassword
					keystorePassword: $keystorePassword
				}
			) {
				id
				keyAlias
				keyPassword
				keystore
				keystorePassword
			}
		}
	}
`

func (service *service) Create(input CreateData) (*Data, error) {
	variables := map[string]any{
		"accountId":        input.AccountId,
		"keystoreBase64":   input.KeystoreBase64,
		"keyAlias":         input.KeyAlias,
		"keyPassword":      input.KeyPassword,
		"keystorePassword": input.KeystorePassword,
	}

	var response createResponse
	err := service.graphql.Query(createQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	return &response.AndroidKeystore.Data, nil
}
