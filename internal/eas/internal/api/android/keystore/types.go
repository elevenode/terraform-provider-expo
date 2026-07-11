package keystore

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type Data struct {
	Id               string `json:"id"`
	KeyAlias         string `json:"keyAlias"`
	KeyPassword      string `json:"keyPassword"`
	Keystore         string `json:"keystore"`
	KeystorePassword string `json:"keystorePassword"`
}

type CreateData struct {
	AccountId        string
	KeystoreBase64   string
	KeyAlias         string
	KeyPassword      string
	KeystorePassword string
}

type Service interface {
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
