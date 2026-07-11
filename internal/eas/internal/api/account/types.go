package account

import "terraform-provider-eas/internal/eas/internal/graphql"

type Data struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type AccountService interface {
	GetByName(name string) (*Data, error)
}

type accountService struct {
	graphql graphql.GraphQL
}

var _ AccountService = (*accountService)(nil)

func NewAccountService(graphql graphql.GraphQL) AccountService {
	return &accountService{graphql: graphql}
}
