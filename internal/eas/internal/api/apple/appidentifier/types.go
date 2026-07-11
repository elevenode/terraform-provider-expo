package appidentifier

import "terraform-provider-eas/internal/eas/internal/graphql"

type Data struct {
	Id         string `json:"id"`
	Identifier string `json:"bundleIdentifier"`
}

type CreateData struct {
	AccountId  string
	Identifier string
}

type GetByIdentifierData struct {
	Identifier string
	AccountId  string
}

type Service interface {
	Create(data CreateData) (*Data, error)
	GetByIdentifier(GetByIdentifierData) (*Data, error)
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
