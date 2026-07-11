package appcredentials

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/utils"
)

func TestDelete(t *testing.T) {
	id := "test-id"
	expectedVariables := map[string]any{"id": id}

	config := utils.TestConfig[string, any, any, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Delete",
		Input:              &id,
		MockResponse:       nil,
		ExpectedQuery:      deleteQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       nil,
	}
	utils.Test(t, config)
}
