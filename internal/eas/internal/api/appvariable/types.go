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

const (
	TypeString     = "STRING"
	TypeFileBase64 = "FILE_BASE64"
)

type CreateData struct {
	AppId        string
	Name         string
	Value        string
	Visibility   string
	Environments []string
	// STRING (default) or FILE_BASE64. For FILE_BASE64, Value is the base64-encoded file content.
	Type string
	// File name of a FILE_BASE64 variable
	FileName string
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
