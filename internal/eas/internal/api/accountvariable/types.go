package accountvariable

import (
	"github.com/fintreal/eas-sdk-go/internal/api/environmentvariable"
	"github.com/fintreal/eas-sdk-go/internal/graphql"
)

type Data = environmentvariable.Data
type UpdateData = environmentvariable.UpdateData

type GetData struct {
	Id        string
	AccountId string
}

type CreateData struct {
	AccountId    string
	Name         string
	Value        string
	Visibility   string
	Environments []string
}

type Service interface {
	Get(GetData) (*Data, error)
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
