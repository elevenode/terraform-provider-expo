package googleserviceaccountkey

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/graphql"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	input := &CreateData{
		AccountId: "test-account-id",
		KeyJson:   `{"type":"service_account","project_id":"test-project","client_email":"sa@test-project.iam.gserviceaccount.com"}`,
	}

	expectedData := &Data{
		Id:                   "test-id",
		ProjectIdentifier:    "test-project",
		ClientEmail:          "sa@test-project.iam.gserviceaccount.com",
		ClientIdentifier:     "test-client-identifier",
		PrivateKeyIdentifier: "test-private-key-identifier",
	}

	expectedVariables := map[string]any{
		"accountId": input.AccountId,
		"jsonKey": map[string]any{
			"type":         "service_account",
			"project_id":   "test-project",
			"client_email": "sa@test-project.iam.gserviceaccount.com",
		},
	}

	mockResponse := createResponse{
		GoogleServiceAccountKey: createGoogleServiceAccountKey{Data: *expectedData},
	}

	config := utils.TestConfig[CreateData, Data, createResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Create",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      createQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}

func TestCreateInvalidJson(t *testing.T) {
	service := NewService(graphql.NewGraphQLMock(createResponse{}))

	data, err := service.Create(CreateData{AccountId: "test-account-id", KeyJson: "not json"})

	assert.Nil(t, data)
	assert.ErrorContains(t, err, "not valid JSON")
}

func TestGet(t *testing.T) {
	id := "test-id"
	expectedData := &Data{
		Id:                   id,
		ProjectIdentifier:    "test-project",
		ClientEmail:          "test-client-email",
		ClientIdentifier:     "test-client-identifier",
		PrivateKeyIdentifier: "test-private-key-identifier",
	}

	mockResponse := getByIdResponse{
		GoogleServiceAccountKey: getByIdQuery{Data: *expectedData},
	}

	config := utils.TestConfig[string, Data, getByIdResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Get",
		Input:              &id,
		MockResponse:       mockResponse,
		ExpectedQuery:      getByIdQueryString,
		ExpectedVariables:  map[string]any{"id": id},
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}

func TestDelete(t *testing.T) {
	id := "test-id"

	config := utils.TestConfig[string, any, any, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Delete",
		Input:              &id,
		MockResponse:       nil,
		ExpectedQuery:      deleteQuery,
		ExpectedVariables:  map[string]any{"id": id},
		ExpectedData:       nil,
	}
	utils.Test(t, config)
}

func TestFindByProjectIdentifier(t *testing.T) {
	keys := []Data{
		{Id: "a", ProjectIdentifier: "p1", ClientEmail: "one@p1"},
		{Id: "b", ProjectIdentifier: "p1", ClientEmail: "two@p1"},
		{Id: "c", ProjectIdentifier: "p2", ClientEmail: "one@p2"},
	}

	match, err := findByProjectIdentifier(keys, "p2", "")
	assert.NoError(t, err)
	assert.Equal(t, "c", match.Id)

	match, err = findByProjectIdentifier(keys, "p1", "two@p1")
	assert.NoError(t, err)
	assert.Equal(t, "b", match.Id)

	_, err = findByProjectIdentifier(keys, "p1", "")
	assert.ErrorContains(t, err, "found 2")

	_, err = findByProjectIdentifier(keys, "missing", "")
	assert.ErrorContains(t, err, "couldn't find")
}
