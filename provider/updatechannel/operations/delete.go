package operations

import (
	"context"

	"github.com/elevenode/terraform-provider-expo/internal/client"
	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const emptyBranchMapping = `{"version":0,"data":[]}`

func Delete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*client.EASClient)
	id := d.Get("id").(string)

	var diags diag.Diagnostics

	// EAS rejects deleting a channel that still points at a branch, so the
	// mapping is cleared first.
	if _, err := client.UpdateChannel.Update(eas.UpdateUpdateChannelData{
		Id:            id,
		BranchMapping: emptyBranchMapping,
	}); err != nil {
		return diag.FromErr(err)
	}

	if _, err := client.UpdateChannel.Delete(id); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}
