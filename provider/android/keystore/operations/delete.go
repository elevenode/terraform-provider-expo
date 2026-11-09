package operations

import (
	"context"

	"github.com/elevenode/terraform-provider-expo/internal/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Delete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*client.EASClient)

	if _, err := client.Android.Keystore.Delete(d.Id()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
