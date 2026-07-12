package utils

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/graphql"
	"github.com/stretchr/testify/assert"
)

type TestConfig[Input, ExpectedData, MockResponse, Service any] struct {
	Input              *Input
	ExpectedQuery      string
	ExpectedVariables  map[string]any
	ExpectedData       *ExpectedData
	MockResponse       MockResponse
	NewServiceFunction func(graphql.GraphQL) Service
	FunctionUnderTest  string
}

func Test[Input, ExpectedData, MockResponse, Service any](t *testing.T, config TestConfig[Input, ExpectedData, MockResponse, Service]) {
	graphQLMock := graphql.NewGraphQLMock(config.MockResponse)

	service := config.NewServiceFunction(graphQLMock)

	actualData, actualError := callMethod[Input, ExpectedData](service, config.FunctionUnderTest, config.Input)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, config.ExpectedQuery, actualQuery)
	assert.Equal(t, config.ExpectedVariables, actualVariables)
	assert.Equal(t, config.ExpectedData, actualData)
	assert.Equal(t, nil, actualError)
}

func callMethod[Input, Data any](service any, methodName string, input *Input) (*Data, error) {
	val := reflect.ValueOf(service)
	method := val.MethodByName(methodName)
	if !method.IsValid() {
		panic(fmt.Sprintf("Method %s not found", method))
	}

	var results []reflect.Value
	if input == nil {
		results = method.Call([]reflect.Value{})
	} else {
		results = method.Call([]reflect.Value{reflect.ValueOf(input).Elem()})
	}

	var resultPtr *Data
	if results[0].IsValid() && !results[0].IsNil() {
		result := results[0].Interface()
		resultPtr = result.(*Data)
	}

	var err error
	if !results[1].IsNil() {
		err = results[1].Interface().(error)
	}

	return resultPtr, err
}
