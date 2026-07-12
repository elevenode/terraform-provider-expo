package appbuildcredentials

import "github.com/elevenode/terraform-provider-expo/internal/eas/internal/graphql"

type objWithId struct {
	Id string `json:"id"`
}

type data struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Keystore objWithId `json:"androidKeystore"`
}

type Data struct {
	Id               string
	Name             string
	KeystoreId       string
	AppCredentialsId string
}

type GetData struct {
	Id               string
	AppId            string
	AppCredentialsId string
}

type CreateData struct {
	Name             string
	KeystoreId       string
	AppCredentialsId string
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
