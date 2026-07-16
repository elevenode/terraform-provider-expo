package updatebranch

import "github.com/elevenode/terraform-provider-expo/internal/eas/internal/graphql"

type Data struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateData struct {
	AppId string
	Name  string
}

type GetByNameData struct {
	AppId string
	Name  string
}

type Service interface {
	Create(data CreateData) (*Data, error)
	GetByName(data GetByNameData) (*Data, error)
	Delete(id string) (*any, error)
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
