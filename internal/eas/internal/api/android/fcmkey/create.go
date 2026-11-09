package fcmkey

type androidAppCredentials struct {
	GoogleServiceAccountKeyData googleServiceAccountKeyForFcmV1 `json:"createFcmV1Credential"`
}

type googleServiceAccountKeyForFcmV1 struct {
	Data Data `json:"googleServiceAccountKeyForFcmV1"`
}

type createResponse struct {
	AndroidAppCredentials androidAppCredentials `json:"androidAppCredentials"`
}

const createQuery = `
	mutation ($appCredentialsId: String!, $keyJson: String!) {
    androidAppCredentials {
      createFcmV1Credential(
        androidAppCredentialsId: $appCredentialsId
        credential: $keyJson
      ) {
        googleServiceAccountKeyForFcmV1 {
          keyJson
          id
        }
      }
    }
	}
`

func (s *service) Create(createData CreateData) (*Data, error) {
	variables := map[string]any{
		"keyJson":          createData.KeyJson,
		"appCredentialsId": createData.AppCredentialsId,
	}

	var response createResponse

	err := s.graphql.Query(createQuery, variables, &response)
	if err != nil {
		return nil, err
	}

	return &response.AndroidAppCredentials.GoogleServiceAccountKeyData.Data, nil
}
