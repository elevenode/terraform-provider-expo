package keystore

import (
	"github.com/elevenode/terraform-provider-expo/provider/android/keystore/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Description: "Uploads an Android upload/signing keystore to the EAS account. Reference its `id` from `expo_android_app_credentials.build_credentials.keystore_id`. " +
			"EAS has no API to read a keystore back by id, so changes made outside Terraform are not detected; every argument forces a new keystore.",
		CreateContext: operations.Create,
		ReadContext:   operations.Read,
		DeleteContext: operations.Delete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "EAS Id of the Android Keystore",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"keystore_base64": {
				Description: "Base64-encoded keystore file (`.jks` / `.keystore`), e.g. `filebase64(\"release.jks\")`.",
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				ForceNew:    true,
			},
			"keystore_password": {
				Description: "Password of the keystore.",
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				ForceNew:    true,
			},
			"key_alias": {
				Description: "Alias of the signing key inside the keystore.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"key_password": {
				Description: "Password of the signing key. Omit if it is the same as the keystore password.",
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				ForceNew:    true,
			},
			"md5_certificate_fingerprint": {
				Description: "MD5 fingerprint of the signing certificate.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"sha1_certificate_fingerprint": {
				Description: "SHA-1 fingerprint of the signing certificate (e.g. for registering the app in Firebase).",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"sha256_certificate_fingerprint": {
				Description: "SHA-256 fingerprint of the signing certificate.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}
