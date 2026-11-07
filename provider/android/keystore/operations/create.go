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

	input := eas.CreateAndroidKeystoreData{
		AccountId:        client.AccountId,
		KeystoreBase64:   d.Get("keystore_base64").(string),
		KeystorePassword: d.Get("keystore_password").(string),
		KeyAlias:         d.Get("key_alias").(string),
		KeyPassword:      d.Get("key_password").(string),
	}

	data, err := client.Android.Keystore.Create(input)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.Id)

	fields := map[string]string{
		"md5_certificate_fingerprint":    data.Md5CertificateFingerprint,
		"sha1_certificate_fingerprint":   data.Sha1CertificateFingerprint,
		"sha256_certificate_fingerprint": data.Sha256CertificateFingerprint,
	}
	for key, value := range fields {
		if err := d.Set(key, value); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	return diags
}
