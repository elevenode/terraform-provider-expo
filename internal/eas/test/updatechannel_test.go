package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/elevenode/terraform-provider-expo/internal/eas/test/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func alwaysTrueBranchMapping(branchId string) string {
	return fmt.Sprintf(`{"version":0,"data":[{"branchId":%q,"branchMappingLogic":"true"}]}`, branchId)
}

// awaitChannelDeletion mirrors the provider's Delete: schedule, then poll the
// background job, since the branch cannot be removed until the channel is gone.
func awaitChannelDeletion(t *testing.T, channelId string) {
	t.Helper()

	receipt, err := utils.Client.UpdateChannel.Delete(channelId)
	require.NoError(t, err)

	for i := 0; !receipt.Done(); i++ {
		require.Less(t, i, 60, "timed out waiting for channel deletion")
		time.Sleep(2 * time.Second)

		receipt, err = utils.Client.BackgroundJob.GetById(receipt.Id)
		require.NoError(t, err)
	}

	assert.Equal(t, eas.BackgroundJobStateSuccess, receipt.State)
}

func TestUpdateChannelCreateGetByNameUpdateAndDelete(t *testing.T) {
	name := utils.GenerateRandomString(10)

	branch, err := utils.Client.UpdateBranch.Create(eas.CreateUpdateBranchData{
		AppId: utils.ImmutableAppId,
		Name:  name,
	})
	require.NoError(t, err)

	// A second branch to remap onto: EAS rejects an empty branch mapping, so an
	// update has to point somewhere real.
	otherBranch, err := utils.Client.UpdateBranch.Create(eas.CreateUpdateBranchData{
		AppId: utils.ImmutableAppId,
		Name:  utils.GenerateRandomString(10),
	})
	require.NoError(t, err)

	branchMapping := alwaysTrueBranchMapping(branch.Id)

	createdData, err := utils.Client.UpdateChannel.Create(eas.CreateUpdateChannelData{
		AppId:         utils.ImmutableAppId,
		Name:          name,
		BranchMapping: branchMapping,
	})

	require.NoError(t, err)
	assert.Equal(t, name, createdData.Name)
	assert.NotEmpty(t, createdData.Id)
	assert.JSONEq(t, branchMapping, createdData.BranchMapping)

	fetchedData, err := utils.Client.UpdateChannel.GetByName(eas.GetByNameUpdateChannelData{
		AppId: utils.ImmutableAppId,
		Name:  name,
	})

	require.NoError(t, err)
	assert.Equal(t, createdData.Id, fetchedData.Id)
	assert.JSONEq(t, branchMapping, fetchedData.BranchMapping)

	remapped := alwaysTrueBranchMapping(otherBranch.Id)

	updatedData, err := utils.Client.UpdateChannel.Update(eas.UpdateUpdateChannelData{
		Id:            createdData.Id,
		BranchMapping: remapped,
	})

	require.NoError(t, err)
	assert.Equal(t, createdData.Id, updatedData.Id)
	assert.JSONEq(t, remapped, updatedData.BranchMapping)

	awaitChannelDeletion(t, createdData.Id)

	_, err = utils.Client.UpdateBranch.Delete(branch.Id)
	assert.NoError(t, err)

	_, err = utils.Client.UpdateBranch.Delete(otherBranch.Id)
	assert.NoError(t, err)
}
