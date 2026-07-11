package app

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type Data struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type CreateData struct {
	Name      string
	Slug      string
	AccountId string
}

type UpdateData struct {
	Id   string
	Name string
}

type Service interface {
	Create(data CreateData) (*Data, error)
	Update(data UpdateData) (*Data, error)
	Get(id string) (*Data, error)
	GetByFullName(fullName string) (*Data, error)
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
