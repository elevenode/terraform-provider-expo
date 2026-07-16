package updatechannel

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

var mockDeleteResponse any

func TestDelete(t *testing.T) {
	id := "test-id"
	expectedVariables := map[string]any{"channelId": id}

	config := utils.TestConfig[string, any, any, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Delete",
		Input:              &id,
		MockResponse:       mockDeleteResponse,
		ExpectedQuery:      deleteQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       nil,
	}
	utils.Test(t, config)
}
