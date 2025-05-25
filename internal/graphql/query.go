package graphql

import (
	"context"

	"github.com/machinebox/graphql"
)

// Executes the query and unmarshals the response from the data field into the response object.
//
// @query - GraphQL query/mutation string
//
// @variables - GraphQL variables in a map
//
// @response - response object reference
func (graphQL *graphQL) Query(query string, variables map[string]any, response any) error {
	request := graphql.NewRequest(query)

	request.Header.Set("Authorization", "Bearer "+graphQL.token)

	for key, value := range variables {
		request.Var(key, value)
	}

	ctx := context.Background()

	err := graphQL.client.Run(ctx, request, &response)

	return err
}
