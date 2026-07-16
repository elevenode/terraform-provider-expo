package updatechannel

import (
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/backgroundjob"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/graphql"
)

type Data struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	// BranchMapping is a stringified JSON document describing which branches
	// the channel routes to. See https://docs.expo.dev/eas-update/channel-surfing/
	BranchMapping string `json:"branchMapping"`
}

type CreateData struct {
	AppId         string
	Name          string
	BranchMapping string
}

type UpdateData struct {
	Id            string
	BranchMapping string
}

type GetByNameData struct {
	AppId string
	Name  string
}

type Service interface {
	Create(data CreateData) (*Data, error)
	Update(data UpdateData) (*Data, error)
	GetByName(data GetByNameData) (*Data, error)
	Delete(id string) (*backgroundjob.Data, error)
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
