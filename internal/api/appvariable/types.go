package appvariable

import (
	"github.com/fintreal/eas-sdk-go/internal/graphql"
)

type Data struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
	// PUBLIC, SENSITIVE, SECRET
	Visibility string `json:"visibility"`
	// development, preview, production
	Environments []string `json:"environments"`
}

type GetByNameData struct {
	Name  string
	AppId string
}

type GetData struct {
	Id    string
	AppId string
}

type CreateData struct {
	AppId        string
	Name         string
	Value        string
	Visibility   string
	Environments []string
}

type UpdateData struct {
	Id           string
	Name         string
	Value        string
	Visibility   string
	Environments []string
}

type Service interface {
	Get(GetData) (*Data, error)
	GetByName(GetByNameData) (*Data, error)
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
