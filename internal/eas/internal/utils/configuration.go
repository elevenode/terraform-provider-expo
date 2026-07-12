package utils

import "github.com/elevenode/terraform-provider-expo/internal/eas/internal/graphql"

type Configuration[InputType any, OutputType any, ResponseType any] struct {
	Query        string
	InputMapper  func(InputType) map[string]any
	OutputMapper func(ResponseType, InputType) *OutputType
}

func ConfigurationFunc[InputType, OutputType, ResponseType any](config Configuration[InputType, OutputType, ResponseType], graphql graphql.GraphQL) func(InputType) (*OutputType, error) {
	return func(input InputType) (*OutputType, error) {
		variables := config.InputMapper(input)

		var response ResponseType

		err := graphql.Query(config.Query, variables, &response)

		if err != nil {
			return nil, err
		}

		return config.OutputMapper(response, input), nil
	}
}
