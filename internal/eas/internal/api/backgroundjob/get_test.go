package backgroundjob

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestGet(t *testing.T) {
	id := "test-receipt-id"
	expectedData := &Data{
		Id:         id,
		State:      StateSuccess,
		Tries:      1,
		WillRetry:  false,
		ResultType: "VOID",
	}
	expectedVariables := map[string]any{"id": id}

	mockResponse := getResponse{BackgroundJobReceipt: &backgroundJobReceipt{Data: expectedData}}

	config := utils.TestConfig[string, Data, getResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Get",
		Input:              &id,
		MockResponse:       mockResponse,
		ExpectedQuery:      getQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
