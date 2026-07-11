package appcredentials

import (
	"github.com/fintreal/eas-sdk-go/internal/api/apple/appbuildcredentials"
	"github.com/fintreal/eas-sdk-go/internal/graphql"
)

type objWithId struct {
	Id string `json:"id"`
}

type buildCredentials struct {
	Id                  string    `json:"id"`
	DistributionType    string    `json:"iosDistributionType"`
	ProvisioningProfile objWithId `json:"provisioningProfile"`
	Certificate         objWithId `json:"distributionCertificate"`
	AppCredentials      objWithId `json:"iosAppCredentials"`
}

type data struct {
	Id               string             `json:"id"`
	App              objWithId          `json:"app"`
	AppIdentifier    objWithId          `json:"appleAppIdentifier"`
	AppStoreApiKey   *objWithId         `json:"appStoreConnectApiKeyForSubmissions"`
	PushKey          *objWithId         `json:"pushKey"`
	BuildCredentials []buildCredentials `json:"iosAppBuildCredentialsArray"`
}

type Data struct {
	Id               string
	AppIdentifierId  string
	AppId            string
	AppStoreApiKeyId *string
	PushKeyId        *string
	BuildCredentials []appbuildcredentials.Data
}

type GetData struct {
	Id    string
	AppId string
}

type CreateData struct {
	AppIdentifierId  string
	AppId            string
	AppStoreApiKeyId *string
	PushKeyId        *string
}

type UpdateData struct {
	Id               string
	AppStoreApiKeyId *string
	PushKeyId        *string
}

type Service interface {
	Get(GetData) (*Data, error)
	Create(CreateData) (*Data, error)
	Update(UpdateData) (*Data, error)
	Delete(string) (*any, error)
}

type service struct {
	graphql graphql.GraphQL
}

var _ Service = (*service)(nil)

func NewService(graphql graphql.GraphQL) Service {
	return &service{graphql: graphql}
}
