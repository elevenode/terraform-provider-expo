package certificate

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/utils"
)

func TestGetBySerialNumber(t *testing.T) {
	input := &GetBySerialNumberData{
		AccountId:    "test-account-id",
		SerialNumber: "test-serial-number",
	}
	expectedData := &Data{
		Id:           "test-id",
		SerialNumber: input.SerialNumber,
	}
	expectedVariables := map[string]any{
		"accountId":    input.AccountId,
		"serialNumber": input.SerialNumber,
	}

	mockResponse := utils.AccountResponse[getResponse]{
		Account: utils.Account[getResponse]{
			ById: getResponse{
				Data: []Data{*expectedData},
			},
		},
	}

	config := utils.TestConfig[GetBySerialNumberData, Data, utils.AccountResponse[getResponse], Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "GetBySerialNumber",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      getQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)

	// graphQLMock := graphql.NewGraphQLMock(mockResponse)

	// service := NewCertificateService(graphQLMock)
	// actualResponse, actualErr := service.GetBySerialNumber()

	// actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	// actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	// assert.Equal(t, getQuery, actualQuery)
	// assert.Equal(t, expectedVariables, actualVariables)
	// assert.Equal(t, expectedData, actualResponse)
	// assert.Equal(t, nil, actualErr)
}
