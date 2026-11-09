package operations

import (
	"context"

	"github.com/elevenode/terraform-provider-expo/internal/client"
	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Create(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*client.EASClient)

	input := eas.CreateAndroidAppCredentialsData{
		AppId:                     d.Get("app_id").(string),
		Identifier:                d.Get("identifier").(string),
		GoogleServiceAccountKeyId: d.Get("google_service_account_key_id").(string),
	}
	data, err := client.Android.AppCredentials.Create(input)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.Id)

	if err := d.Set("id", data.Id); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("app_id", data.AppId); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("identifier", data.Identifier); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("google_service_account_key_id", data.GoogleServiceAccountKeyId); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if fcmKey := d.Get("fcm_key").(string); fcmKey != "" {
		fcmKeyInput := eas.CreateFCMKey{
			KeyJson:          fcmKey,
			AppCredentialsId: data.Id,
		}

		dataFcmKey, err := client.Android.FCMKey.Create(fcmKeyInput)
		if err != nil {
			return append(diags, diag.FromErr(err)...)
		}

		if err := d.Set("fcm_key", dataFcmKey.KeyJson); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	buildCredentials, err := createBuildCredentials(client, data.Id, d.Get("build_credentials").([]any))
	if err != nil {
		return append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("build_credentials", buildCredentials); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}
