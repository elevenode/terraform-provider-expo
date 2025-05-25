package graphql

import (
	"github.com/fintreal/eas-sdk-go/internal/graphql/machinebox"
	"github.com/machinebox/graphql"
)

type VariableValue interface {
	~int | ~string | ~[]string
}

type GraphQL interface {
	Query(query string, variables map[string]any, response any) error
}

type graphQL struct {
	client machinebox.MachineboxGraphqlClient
	token  string
}

var _ GraphQL = (*graphQL)(nil)

func NewGraphQL(token string) GraphQL {
	client := graphql.NewClient("https://api.expo.dev/graphql")
	return &graphQL{
		client: client,
		token:  token,
	}
}
