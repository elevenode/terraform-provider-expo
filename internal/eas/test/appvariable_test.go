package test

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/appvariable"
	"github.com/elevenode/terraform-provider-expo/internal/eas/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestAppEnvironmentVariableGet(t *testing.T) {
	input := eas.GetAppVariableData{
		Id:    utils.ImmutableAppVariableId,
		AppId: utils.ImmutableAppId,
	}

	actualData, actualErr := utils.Client.AppVariable.Get(input)

	expectedData := &appvariable.Data{
		Id:           utils.ImmutableAppVariableId,
		Name:         "TEST_ENVIRONMENT_VARIABLE",
		Value:        "VALUE",
		Visibility:   "PUBLIC",
		Environments: []string{"preview"},
	}

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestAppEnvironmentVariableGetByName(t *testing.T) {
	input := eas.GetByNameAppVariableData{
		Name:  "TEST_ENVIRONMENT_VARIABLE",
		AppId: utils.ImmutableAppId,
	}

	actualData, actualErr := utils.Client.AppVariable.GetByName(input)

	expectedData := &appvariable.Data{
		Id:           utils.ImmutableAppVariableId,
		Name:         input.Name,
		Value:        "VALUE",
		Visibility:   "PUBLIC",
		Environments: []string{"preview"},
	}

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestAppEnvironmentVariableUpdate(t *testing.T) {
	updateData := appvariable.UpdateData{
		Id:           utils.MutableAppVariableId,
		Name:         utils.GenerateRandomString(10),
		Value:        utils.GenerateRandomString(10),
		Visibility:   "PUBLIC",
		Environments: []string{"production"},
	}

	actualData, actualErr := utils.Client.AppVariable.Update(updateData)

	assert.Equal(t, actualData.Id, updateData.Id)
	assert.Equal(t, actualData.Name, updateData.Name)
	assert.Equal(t, actualData.Value, updateData.Value)
	assert.Equal(t, actualData.Visibility, updateData.Visibility)
	assert.Equal(t, actualData.Environments, updateData.Environments)
	assert.Equal(t, nil, actualErr)
}

func TestAppEnvironmentVariableCreateAndDelete(t *testing.T) {
	inputData := appvariable.CreateData{
		AppId:        utils.ImmutableAppId,
		Name:         utils.GenerateRandomString(10),
		Value:        utils.GenerateRandomString(10),
		Visibility:   "PUBLIC",
		Environments: []string{"development", "preview"},
	}

	actualData, actualErr := utils.Client.AppVariable.Create(inputData)

	assert.Equal(t, inputData.Name, actualData.Name)
	assert.Equal(t, inputData.Value, actualData.Value)
	assert.Equal(t, inputData.Visibility, actualData.Visibility)
	assert.Equal(t, inputData.Environments, actualData.Environments)
	assert.Equal(t, nil, actualErr)

	_, actualErr = utils.Client.AppVariable.Delete(actualData.Id)

	assert.Equal(t, nil, actualErr)
}
