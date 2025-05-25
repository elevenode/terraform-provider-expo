package appstoreapikey

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type Data struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	IssuerIdentifier string `json:"issuerIdentifier"`
	Identifier       string `json:"keyIdentifier"`
}

type GeyByIdentifierData struct {
	Identifier string
	AccountId  string
}

type Service interface {
	GetByIdentifier(GeyByIdentifierData) (*Data, error)
}

type service struct {
	graphql graphql.GraphQL
}

var _ Service = (*service)(nil)

func NewService(graphql graphql.GraphQL) Service {
	return &service{
		graphql: graphql,
	}
}
