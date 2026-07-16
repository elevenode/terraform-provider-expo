package test

import (
	"fmt"
	"testing"

	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/elevenode/terraform-provider-expo/internal/eas/test/utils"
	"github.com/stretchr/testify/assert"
)

const emptyBranchMapping = `{"version":0,"data":[]}`

func alwaysTrueBranchMapping(branchId string) string {
	return fmt.Sprintf(`{"version":0,"data":[{"branchId":%q,"branchMappingLogic":"true"}]}`, branchId)
}

func TestUpdateChannelCreateGetByNameUpdateAndDelete(t *testing.T) {
	name := utils.GenerateRandomString(10)

	branch, actualErr := utils.Client.UpdateBranch.Create(eas.CreateUpdateBranchData{
		AppId: utils.ImmutableAppId,
		Name:  name,
	})
	assert.Equal(t, nil, actualErr)

	branchMapping := alwaysTrueBranchMapping(branch.Id)

	createdData, actualErr := utils.Client.UpdateChannel.Create(eas.CreateUpdateChannelData{
		AppId:         utils.ImmutableAppId,
		Name:          name,
		BranchMapping: branchMapping,
	})

	assert.Equal(t, nil, actualErr)
	assert.Equal(t, name, createdData.Name)
	assert.NotEmpty(t, createdData.Id)
	assert.JSONEq(t, branchMapping, createdData.BranchMapping)

	fetchedData, actualErr := utils.Client.UpdateChannel.GetByName(eas.GetByNameUpdateChannelData{
		AppId: utils.ImmutableAppId,
		Name:  name,
	})

	assert.Equal(t, nil, actualErr)
	assert.Equal(t, createdData.Id, fetchedData.Id)
	assert.JSONEq(t, branchMapping, fetchedData.BranchMapping)

	// Clearing the mapping is what makes the channel deletable, so this asserts
	// the precondition the provider's Delete relies on.
	updatedData, actualErr := utils.Client.UpdateChannel.Update(eas.UpdateUpdateChannelData{
		Id:            createdData.Id,
		BranchMapping: emptyBranchMapping,
	})

	assert.Equal(t, nil, actualErr)
	assert.Equal(t, createdData.Id, updatedData.Id)
	assert.JSONEq(t, emptyBranchMapping, updatedData.BranchMapping)

	_, actualErr = utils.Client.UpdateChannel.Delete(createdData.Id)
	assert.Equal(t, nil, actualErr)

	_, actualErr = utils.Client.UpdateBranch.Delete(branch.Id)
	assert.Equal(t, nil, actualErr)
}
