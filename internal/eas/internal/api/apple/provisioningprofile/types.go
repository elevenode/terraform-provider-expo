package provisioningprofile

import "github.com/elevenode/terraform-provider-expo/internal/eas/internal/graphql"

type appleAppIdentifier struct {
	Id string `json:"id"`
}

type data struct {
	Id                 string             `json:"id"`
	Base64             string             `json:"provisioningProfile"`
	AppleAppIdentifier appleAppIdentifier `json:"appleAppIdentifier"`
}

type Data struct {
	Id                    string
	Base64                string
	AppBundleIdentifierId string
}

type CreateData struct {
	Base64                string
	AppBundleIdentifierId string
	AccountId             string
}

type GetData struct {
	Id        string
	AccountId string
}

type Service interface {
	Get(GetData) (*Data, error)
	Create(data CreateData) (*Data, error)
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
