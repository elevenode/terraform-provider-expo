package updatechannel

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/backgroundjob"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestDelete(t *testing.T) {
	id := "test-id"
	expectedVariables := map[string]any{"channelId": id}
	expectedData := &backgroundjob.Data{
		Id:         "test-receipt-id",
		State:      backgroundjob.StateQueued,
		ResultType: "UpdateChannel",
	}

	mockResponse := deleteResponse{UpdateChannel: scheduleUpdateChannelDeletion{Data: expectedData}}

	config := utils.TestConfig[string, backgroundjob.Data, deleteResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Delete",
		Input:              &id,
		MockResponse:       mockResponse,
		ExpectedQuery:      deleteQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
