package operations

import (
	"context"

	"github.com/elevenode/terraform-provider-expo/internal/client"
	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*client.EASClient)

	data, err := client.Android.GoogleServiceAccountKey.Create(eas.CreateGoogleServiceAccountKeyData{
		AccountId: client.AccountId,
		KeyJson:   d.Get("key_json").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.Id)

	return setResourceData(d, data)
}

func ResourceRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*client.EASClient)

	data, err := client.Android.GoogleServiceAccountKey.Get(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return setResourceData(d, data)
}

func ResourceDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*client.EASClient)

	if _, err := client.Android.GoogleServiceAccountKey.Delete(d.Id()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func setResourceData(d *schema.ResourceData, data *eas.GoogleServiceAccountKeyData) diag.Diagnostics {
	var diags diag.Diagnostics

	fields := map[string]string{
		"project_identifier":     data.ProjectIdentifier,
		"client_email":           data.ClientEmail,
		"client_identifier":      data.ClientIdentifier,
		"private_key_identifier": data.PrivateKeyIdentifier,
	}
	for key, value := range fields {
		if err := d.Set(key, value); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	return diags
}
