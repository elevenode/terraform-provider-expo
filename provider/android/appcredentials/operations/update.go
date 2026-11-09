package operations

import (
	"context"

	"github.com/elevenode/terraform-provider-expo/internal/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Update(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	if !d.HasChange("build_credentials") {
		return diags
	}

	client := m.(*client.EASClient)

	// Build credentials have no update mutation: replace the whole list.
	old, new := d.GetChange("build_credentials")
	for _, b := range old.([]any) {
		buildCredential := b.(map[string]any)
		id, _ := buildCredential["id"].(string)
		if id == "" {
			continue
		}
		if _, err := client.Android.AppBuildCredentials.Delete(id); err != nil {
			return diag.FromErr(err)
		}
	}

	created, err := createBuildCredentials(client, d.Id(), new.([]any))
	if err != nil {
		// Record what was created so the next plan reconciles it.
		_ = d.Set("build_credentials", created)
		return diag.FromErr(err)
	}

	if err := d.Set("build_credentials", created); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	return diags
}
