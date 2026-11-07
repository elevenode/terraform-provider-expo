package googleserviceaccountkey

import "github.com/elevenode/terraform-provider-expo/internal/eas/internal/graphql"

type GetByProjectIdentifierData struct {
	ProjectIdentifier string
	// Optional: narrows the match when several keys share a GCP project.
	ClientEmail string
	AccountId   string
}

type CreateData struct {
	AccountId string
	// JSON contents of the Google service account key file.
	KeyJson string
}

type Data struct {
	Id                   string `json:"id"`
	ProjectIdentifier    string `json:"projectIdentifier"`
	ClientEmail          string `json:"clientEmail"`
	ClientIdentifier     string `json:"clientIdentifier"`
	PrivateKeyIdentifier string `json:"privateKeyIdentifier"`
}

type service struct {
	graphql graphql.GraphQL
}

type Service interface {
	GetByProjectIdentifier(GetByProjectIdentifierData) (*Data, error)
	Get(string) (*Data, error)
	Create(CreateData) (*Data, error)
	Delete(string) (*any, error)
}

var _ Service = (*service)(nil)

func NewService(graphql graphql.GraphQL) Service {
	return &service{
		graphql: graphql,
	}
}
