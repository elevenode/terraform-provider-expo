package appvariable

import (
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/environmentvariable"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/graphql"
)

type Data = environmentvariable.Data
type UpdateData = environmentvariable.UpdateData

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

type Service interface {
	Get(GetData) (*Data, error)
	GetByName(GetByNameData) (*Data, error)
	Create(CreateData) (*Data, error)
	Update(UpdateData) (*Data, error)
	Delete(string) (*any, error)
}

type service struct {
	environmentvariable.BaseService
	graphql graphql.GraphQL
}

var _ Service = (*service)(nil)

func NewService(graphql graphql.GraphQL) Service {
	return &service{
		BaseService: environmentvariable.NewBaseService(graphql),
		graphql:     graphql,
	}
}
