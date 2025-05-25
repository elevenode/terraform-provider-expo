package graphql

import "github.com/stretchr/testify/mock"

type GraphQLMock struct {
	mock.Mock
}

var _ GraphQL = (*GraphQLMock)(nil)

func (m *GraphQLMock) Query(query string, variables map[string]any, response any) error {
	args := m.Called(query, variables, response)
	return args.Error(0)
}

func NewGraphQLMock[T any](response T) *GraphQLMock {
	graphQLMock := GraphQLMock{}
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		*args.Get(2).(*T) = response
	}).Return(nil)
	return &graphQLMock
}
