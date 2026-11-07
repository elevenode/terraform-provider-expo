package googleserviceaccountkey

import (
	"encoding/json"
	"strings"

	"github.com/elevenode/terraform-provider-expo/provider/android/googleserviceaccountkey/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Description: "Uploads a Google service account key (JSON key file) to the EAS account. " +
			"Use its `id` as `expo_android_app_credentials.google_service_account_key_id` for Play Store submissions. " +
			"To reference a key that is already in EAS, use the `expo_google_service_account_key` data source instead.",
		CreateContext: operations.ResourceCreate,
		ReadContext:   operations.ResourceRead,
		DeleteContext: operations.ResourceDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "EAS Id of the Google Service Account Key",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"key_json": {
				Description:      "JSON contents of the Google service account key file, e.g. `file(\"service-account.json\")`.",
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				ForceNew:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: jsonEqual,
			},
			"project_identifier": {
				Description: "GCP project id of the key (`project_id` in the JSON key file)",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"client_email": {
				Description: "Google Service Account Email",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"client_identifier": {
				Description: "Google Service Account Identifier",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"private_key_identifier": {
				Description: "Identifier of the private key (`private_key_id` in the JSON key file)",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

// jsonEqual ignores formatting differences in the key JSON.
func jsonEqual(k, old, new string, d *schema.ResourceData) bool {
	if strings.TrimSpace(old) == "" || strings.TrimSpace(new) == "" {
		return strings.TrimSpace(old) == strings.TrimSpace(new)
	}
	var oldJSON, newJSON any
	if json.Unmarshal([]byte(old), &oldJSON) != nil || json.Unmarshal([]byte(new), &newJSON) != nil {
		return old == new
	}
	oldNorm, _ := json.Marshal(oldJSON)
	newNorm, _ := json.Marshal(newJSON)
	return string(oldNorm) == string(newNorm)
}
