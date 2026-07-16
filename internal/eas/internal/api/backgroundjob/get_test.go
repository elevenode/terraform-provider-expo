package backgroundjob

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
)

func TestGetById(t *testing.T) {
	expectedData := &Data{
		Id:         "test-receipt-id",
		State:      StateSuccess,
		Tries:      1,
		ResultType: "UpdateChannel",
	}

	mockResponse := getResponse{BackgroundJobReceipt: receiptById{Data: expectedData}}

	config := utils.TestConfig[string, Data, getResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "GetById",
		Input:              &expectedData.Id,
		MockResponse:       mockResponse,
		ExpectedQuery:      getByIdQuery,
		ExpectedVariables:  map[string]any{"id": expectedData.Id},
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}

func TestDone(t *testing.T) {
	cases := []struct {
		state     State
		willRetry bool
		want      bool
	}{
		{StateQueued, false, false},
		{StateInProgress, false, false},
		{StateSuccess, false, true},
		{StateFailure, true, false},
		{StateFailure, false, true},
	}

	for _, c := range cases {
		data := &Data{State: c.state, WillRetry: c.willRetry}
		if got := data.Done(); got != c.want {
			t.Errorf("Done() with state %s willRetry %v = %v, want %v", c.state, c.willRetry, got, c.want)
		}
	}
}
