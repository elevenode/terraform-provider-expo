package environmentvariable

import (
	"testing"

	"terraform-provider-eas/internal/eas/internal/utils"
)

var mockDeleteResponse any

func TestDelete(t *testing.T) {
	id := "test-id"
	expectedVariables := map[string]any{"id": id}

	config := utils.TestConfig[string, any, any, BaseService]{
		NewServiceFunction: NewBaseService,
		FunctionUnderTest:  "Delete",
		Input:              &id,
		MockResponse:       mockDeleteResponse,
		ExpectedQuery:      deleteQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       nil,
	}
	utils.Test(t, config)
}
