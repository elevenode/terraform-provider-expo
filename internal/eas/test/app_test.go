package test

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/app"
	"github.com/elevenode/terraform-provider-expo/internal/eas/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestAppGet(t *testing.T) {
	expectedData := &app.Data{
		Id:   utils.ImmutableAppId,
		Name: "Integration Test Fixture",
		Slug: "test-app",
	}

	actualData, actualErr := utils.Client.App.Get(expectedData.Id)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestAppGetByFullName(t *testing.T) {
	expectedData := &app.Data{
		Id:   utils.ImmutableAppId,
		Name: "Integration Test Fixture",
		Slug: "test-app",
	}

	actualData, actualErr := utils.Client.App.GetByFullName("@expo-eas-sdk-go/test-app")

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestAppCreate(t *testing.T) {
	inputData := app.CreateData{
		AccountId: utils.AccountId,
		Name:      utils.GenerateRandomString(10),
		Slug:      utils.GenerateRandomString(10),
	}

	// Create
	expectedData, actualErr := utils.Client.App.Create(inputData)

	assert.Equal(t, nil, actualErr)
	assert.NotNil(t, expectedData)
	if expectedData != nil {
		assert.Equal(t, inputData.Name, expectedData.Name)
		assert.Equal(t, inputData.Slug, expectedData.Slug)
	}
}

func TestAppUpdate(t *testing.T) {
	expectedData := &app.Data{
		Id:   utils.MutableAppId,
		Name: utils.GenerateRandomString(10),
		Slug: "test-app-update",
	}

	updateData := app.UpdateData{
		Id:   expectedData.Id,
		Name: expectedData.Name,
	}

	actualData, actualErr := utils.Client.App.Update(updateData)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
