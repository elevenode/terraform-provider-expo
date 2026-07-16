package test

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/app"
	"github.com/elevenode/terraform-provider-expo/internal/eas/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestAppGet(t *testing.T) {
	expectedData := &app.Data{
		Id:   utils.ImmutableAppId,
		Name: "Test App",
		Slug: "test-app",
	}

	actualData, actualErr := utils.Client.App.Get(expectedData.Id)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestAppGetByFullName(t *testing.T) {
	expectedData := &app.Data{
		Id:   utils.ImmutableAppId,
		Name: "Test App",
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
		// Registered before the assertions so a failure below still can't leak the project.
		t.Cleanup(func() {
			assert.NoError(t, utils.DeleteApp(expectedData.Id))
		})

		assert.Equal(t, inputData.Name, expectedData.Name)
		assert.Equal(t, inputData.Slug, expectedData.Slug)
	}
}

func TestAppDelete(t *testing.T) {
	created, createErr := utils.Client.App.Create(app.CreateData{
		AccountId: utils.AccountId,
		Name:      utils.GenerateRandomString(10),
		Slug:      utils.GenerateRandomString(10),
	})

	assert.Equal(t, nil, createErr)
	assert.NotNil(t, created)
	if created == nil {
		return
	}

	receipt, actualErr := utils.Client.App.Delete(created.Id)

	assert.Equal(t, nil, actualErr)
	assert.NotNil(t, receipt)
	if receipt == nil {
		return
	}
	assert.NotEmpty(t, receipt.Id)

	finished, waitErr := utils.WaitForBackgroundJob(receipt.Id, utils.AppDeletionTimeout)

	assert.NoError(t, waitErr)
	if finished != nil {
		assert.Equal(t, eas.BackgroundJobStateSuccess, finished.State)
	}

	deleted, getErr := utils.Client.App.Get(created.Id)
	assert.True(t, getErr != nil || deleted == nil, "expected app %s to be gone after deletion", created.Id)
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
