package fcmkey

import "terraform-provider-eas/internal/eas/internal/graphql"

type Data struct {
	Id      string `json:"id"`
	KeyJson string `json:"keyJson"`
}

type CreateData struct {
	KeyJson          string
	AccountId        string
	AppCredentialsId string
}

type Service interface {
	Create(CreateData) (*Data, error)
}

type service struct {
	graphql graphql.GraphQL
}

var _ Service = (*service)(nil)

func NewService(graphql graphql.GraphQL) Service {
	return &service{graphql: graphql}
}
