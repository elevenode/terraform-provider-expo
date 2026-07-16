package operations

import (
	"context"

	"github.com/elevenode/terraform-provider-expo/internal/client"
	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Update(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*client.EASClient)

	input := eas.UpdateUpdateChannelData{
		Id:            d.Get("id").(string),
		BranchMapping: d.Get("branch_mapping").(string),
	}

	data, err := client.UpdateChannel.Update(input)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(data.Id)

	var diags diag.Diagnostics
	return diags
}
