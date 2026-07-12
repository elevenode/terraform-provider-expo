package certificate

import "github.com/elevenode/terraform-provider-expo/internal/eas/internal/graphql"

type Data struct {
	Id           string `json:"id"`
	SerialNumber string `json:"serialNumber"`
}

type GetBySerialNumberData struct {
	SerialNumber string
	AccountId    string
}

type Service interface {
	GetBySerialNumber(GetBySerialNumberData) (*Data, error)
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
