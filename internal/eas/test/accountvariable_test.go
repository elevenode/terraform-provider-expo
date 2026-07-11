package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"terraform-provider-eas/internal/eas"
	"terraform-provider-eas/internal/eas/internal/api/accountvariable"
	"terraform-provider-eas/internal/eas/test/utils"
)

func TestAccountEnvironmentVariableCreateAndDelete(t *testing.T) {
	inputData := accountvariable.CreateData{
		AccountId:    utils.AccountId,
		Name:         utils.GenerateRandomString(10),
		Value:        utils.GenerateRandomString(10),
		Visibility:   "PUBLIC",
		Environments: []string{"development", "preview"},
	}

	actualData, actualErr := utils.Client.AccountVariable.Create(inputData)

	assert.Equal(t, nil, actualErr)
	assert.Equal(t, inputData.Name, actualData.Name)
	assert.Equal(t, inputData.Value, actualData.Value)
	assert.Equal(t, inputData.Visibility, actualData.Visibility)
	assert.Equal(t, inputData.Environments, actualData.Environments)
	assert.NotEmpty(t, actualData.Id)
	assert.NotEmpty(t, actualData.CreatedAt)
	assert.NotEmpty(t, actualData.UpdatedAt)

	// Test Get
	getData := eas.GetAccountVariableData{
		Id:        actualData.Id,
		AccountId: utils.AccountId,
	}
	fetchedData, fetchErr := utils.Client.AccountVariable.Get(getData)

	assert.Equal(t, nil, fetchErr)
	assert.Equal(t, actualData.Id, fetchedData.Id)
	assert.Equal(t, actualData.Name, fetchedData.Name)

	// Test Update
	updateData := accountvariable.UpdateData{
		Id:           actualData.Id,
		Name:         utils.GenerateRandomString(10),
		Value:        utils.GenerateRandomString(10),
		Visibility:   "PUBLIC",
		Environments: []string{"production"},
	}

	updatedData, updateErr := utils.Client.AccountVariable.Update(updateData)

	assert.Equal(t, nil, updateErr)
	assert.Equal(t, updateData.Id, updatedData.Id)
	assert.Equal(t, updateData.Name, updatedData.Name)
	assert.Equal(t, updateData.Value, updatedData.Value)
	assert.Equal(t, updateData.Visibility, updatedData.Visibility)
	assert.Equal(t, updateData.Environments, updatedData.Environments)

	// Test Delete
	_, deleteErr := utils.Client.AccountVariable.Delete(actualData.Id)

	assert.Equal(t, nil, deleteErr)
}

func TestAccountEnvironmentVariableGet(t *testing.T) {
	if utils.ImmutableAccountVariableId == "" {
		t.Skip("IMMUTABLE_ACCOUNT_VARIABLE_ID not set")
	}

	input := eas.GetAccountVariableData{
		Id:        utils.ImmutableAccountVariableId,
		AccountId: utils.AccountId,
	}

	actualData, actualErr := utils.Client.AccountVariable.Get(input)

	assert.Equal(t, nil, actualErr)
	assert.Equal(t, utils.ImmutableAccountVariableId, actualData.Id)
	assert.NotEmpty(t, actualData.Name)
}

func TestAccountEnvironmentVariableUpdate(t *testing.T) {
	if utils.MutableAccountVariableId == "" {
		t.Skip("MUTABLE_ACCOUNT_VARIABLE_ID not set")
	}

	updateData := accountvariable.UpdateData{
		Id:           utils.MutableAccountVariableId,
		Name:         utils.GenerateRandomString(10),
		Value:        utils.GenerateRandomString(10),
		Visibility:   "PUBLIC",
		Environments: []string{"production"},
	}

	actualData, actualErr := utils.Client.AccountVariable.Update(updateData)

	assert.Equal(t, nil, actualErr)
	assert.Equal(t, actualData.Id, updateData.Id)
	assert.Equal(t, actualData.Name, updateData.Name)
	assert.Equal(t, actualData.Value, updateData.Value)
	assert.Equal(t, actualData.Visibility, updateData.Visibility)
	assert.Equal(t, actualData.Environments, updateData.Environments)
}
