package appcredentials

import (
	"encoding/json"
	"github.com/elevenode/terraform-provider-expo/provider/android/appcredentials/operations"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Description:   "Android credentials for an EAS app: the Google service account key used for Play Store submissions, the FCM V1 push key, and the keystores used to sign builds.",
		ReadContext:   operations.Read,
		CreateContext: operations.Create,
		DeleteContext: operations.Delete,
		UpdateContext: operations.Update,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "EAS Android App Credential Id",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"app_id": {
				Description: "EAS App Id",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"identifier": {
				Description: "Identifier of the Android App",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"google_service_account_key_id": {
				Description: "EAS Id of the Google Service Account Key used for Play Store submissions. Use the `id` of an `expo_google_service_account_key` resource (to upload a key) or data source (to look up one already in EAS).",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"fcm_key": {
				Description: "JSON contents of the Google service account key used for FCM V1 push notifications.",
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				ForceNew:    true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// If both are empty, they're equal
					if strings.TrimSpace(old) == "" && strings.TrimSpace(new) == "" {
						return true
					}

					// Try to parse as JSON and compare the normalized form
					var oldJSON, newJSON interface{}
					if err := json.Unmarshal([]byte(old), &oldJSON); err == nil {
						if err := json.Unmarshal([]byte(new), &newJSON); err == nil {
							// Re-marshal both to get normalized JSON
							oldNorm, _ := json.Marshal(oldJSON)
							newNorm, _ := json.Marshal(newJSON)
							return string(oldNorm) == string(newNorm)
						}
					}

					// Fallback: normalize whitespace by removing all whitespace and comparing
					oldNormalized := strings.Join(strings.Fields(old), "")
					newNormalized := strings.Join(strings.Fields(new), "")
					return oldNormalized == newNormalized
				},
			},
			"build_credentials": {
				Description: "EAS Android Build Credentials",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "EAS Android Build Credential Id",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"keystore_id": {
							Description: "EAS Id of the Android Keystore, e.g. the `id` of an `expo_android_keystore` resource.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"name": {
							Description: "Name of the Android Build Credential",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
		},
	}
}
