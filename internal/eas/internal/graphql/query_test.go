package graphql

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"terraform-provider-eas/internal/eas/internal/graphql/machinebox"
)

func TestQuery(t *testing.T) {
	client := newTestQueryGrapQLClientMock()
	token := "test-token"
	graphQL := graphQL{client: client, token: token}

	expectedAuthHeader := "Bearer " + token
	expectedQuery := "test-query"
	expectedVariables := map[string]any{"testKey": "test-value"}

	actualError := graphQL.Query(expectedQuery, expectedVariables, nil)
	actualAuthHeader, actualQuery, actualVariables := extractMockCallProperties(client)

	assert.Equal(t, nil, actualError)
	assert.Equal(t, expectedQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedAuthHeader, actualAuthHeader)
}

func newTestQueryGrapQLClientMock() *machinebox.MachineBoxGraphQLClientMock {
	client := machinebox.MachineBoxGraphQLClientMock{}
	client.On("Run", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	return &client
}

func extractMockCallProperties(client *machinebox.MachineBoxGraphQLClientMock) (string, string, any) {
	request := client.Calls[0].Arguments.Get(1).(*graphql.Request)
	requestReflection := reflect.ValueOf(request).Elem()

	authHeader := request.Header.Get("Authorization")
	query := requestReflection.FieldByName("q").String()
	variables := *(*map[string]any)(unsafe.Pointer(requestReflection.FieldByName("vars").UnsafeAddr()))
	return authHeader, query, variables
}
