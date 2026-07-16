package test

import (
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/elevenode/terraform-provider-expo/internal/eas/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestUpdateBranchCreateGetByNameAndDelete(t *testing.T) {
	name := utils.GenerateRandomString(10)

	createdData, actualErr := utils.Client.UpdateBranch.Create(eas.CreateUpdateBranchData{
		AppId: utils.ImmutableAppId,
		Name:  name,
	})

	assert.Equal(t, nil, actualErr)
	assert.Equal(t, name, createdData.Name)
	assert.NotEmpty(t, createdData.Id)

	fetchedData, actualErr := utils.Client.UpdateBranch.GetByName(eas.GetByNameUpdateBranchData{
		AppId: utils.ImmutableAppId,
		Name:  name,
	})

	assert.Equal(t, nil, actualErr)
	assert.Equal(t, createdData, fetchedData)

	_, actualErr = utils.Client.UpdateBranch.Delete(createdData.Id)

	assert.Equal(t, nil, actualErr)
}
