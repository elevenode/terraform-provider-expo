package appcredentials

import (
	"terraform-provider-eas/internal/eas/internal/api/android/appbuildcredentials"
	"terraform-provider-eas/internal/eas/internal/graphql"
)

type objWithId struct {
	Id string `json:"id"`
}

type fcmKey struct {
	Id      string `json:"id"`
	KeyJson string `json:"keyJson"`
}

type buildCredentialsData struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Keystore objWithId `json:"androidKeystore"`
}

type data struct {
	Id                      string                 `json:"id"`
	Identifier              string                 `json:"applicationIdentifier"`
	App                     objWithId              `json:"app"`
	GoogleServiceAccountKey objWithId              `json:"googleServiceAccountKeyForSubmissions"`
	FCMKey                  *fcmKey                `json:"googleServiceAccountKeyForFcmV1"`
	BuildCredentials        []buildCredentialsData `json:"androidAppBuildCredentialsArray"`
}

type Data struct {
	Id                        string
	AppId                     string
	Identifier                string
	GoogleServiceAccountKeyId string
	FCMKey                    *string
	FCMKeyId                  *string
	BuildCredentials          []appbuildcredentials.Data
}

type GetData struct {
	Id    string
	AppId string
}

type CreateData struct {
	AppId                     string
	Identifier                string
	GoogleServiceAccountKeyId string
}

type Service interface {
	Get(GetData) (*Data, error)
	Create(CreateData) (*Data, error)
	Delete(string) (*any, error)
}

type service struct {
	graphql graphql.GraphQL
}

var _ Service = (*service)(nil)

func NewService(graphql graphql.GraphQL) Service {
	return &service{graphql: graphql}
}
