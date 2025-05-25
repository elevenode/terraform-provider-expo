package googleserviceaccountkey

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type GetByProjectIdentifierData struct {
	ProjectIdentifier string
	AccountId         string
}

type Data struct {
	Id                string `json:"id"`
	ProjectIdentifier string `json:"projectIdentifier"`
	ClientEmail       string `json:"clientEmail"`
	ClientIdentifier  string `json:"clientIdentifier"`
}

type service struct {
	graphql graphql.GraphQL
}

type Service interface {
	GetByProjectIdentifier(GetByProjectIdentifierData) (*Data, error)
}

var _ Service = (*service)(nil)

func NewService(graphql graphql.GraphQL) Service {
	return &service{
		graphql: graphql,
	}
}
