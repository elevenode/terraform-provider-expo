package environmentvariable

import (
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/graphql"
)

type Data struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
	// PUBLIC, SENSITIVE, SECRET
	Visibility string `json:"visibility"`
	// development, preview, production
	Environments []string `json:"environments"`
	// STRING or FILE_BASE64
	Type string `json:"type,omitempty"`
	// File name of a FILE_BASE64 variable
	FileName  string `json:"fileName,omitempty"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UpdateData struct {
	Id           string
	Name         string
	Value        string
	Visibility   string
	Environments []string
	// Optional: STRING or FILE_BASE64. Left unchanged when empty.
	Type string
	// Optional: file name of a FILE_BASE64 variable. Left unchanged when empty.
	FileName string
}

type BaseService interface {
	Update(UpdateData) (*Data, error)
	Delete(string) (*any, error)
}

type baseService struct {
	graphql graphql.GraphQL
}

var _ BaseService = (*baseService)(nil)

func NewBaseService(graphql graphql.GraphQL) BaseService {
	return &baseService{graphql: graphql}
}

func (s *baseService) GetGraphQL() graphql.GraphQL {
	return s.graphql
}
