package team

import "terraform-provider-eas/internal/eas/internal/graphql"

type account struct {
	Id string `json:"id"`
}

type data struct {
	Id         string  `json:"id"`
	Name       string  `json:"appleTeamName"`
	Identifier string  `json:"appleTeamIdentifier"`
	Type       string  `json:"appleTeamType"`
	Account    account `json:"account"`
}

type Data struct {
	Id         string
	Name       string
	Identifier string
	Type       string
	AccountId  string
}

type CreateData struct {
	Name       string
	Identifier string
	Type       string
	AccountId  string
}

type UpdateData struct {
	Id   string
	Name string
	Type string
}

type GetByIdentifierData struct {
	Identifier string
	AccountId  string
}

type Service interface {
	Create(data CreateData) (*Data, error)
	Update(data UpdateData) (*Data, error)
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
