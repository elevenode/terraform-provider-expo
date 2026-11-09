package googleserviceaccountkey

import (
	"github.com/elevenode/terraform-provider-expo/provider/android/googleserviceaccountkey/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Looks up a Google service account key that is already uploaded to the EAS account. To upload one with Terraform, use the `expo_google_service_account_key` resource.",
		ReadContext: operations.Read,
		Schema: map[string]*schema.Schema{
			"project_identifier": {
				Description: "GCP project id of the key (`project_id` in the JSON key file)",
				Type:        schema.TypeString,
				Required:    true,
			},
			"id": {
				Description: "EAS Id of the Google Service Account Key",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"client_email": {
				Description: "Google Service Account Email. Set it to pick one key when several keys in EAS belong to the same GCP project.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"client_identifier": {
				Description: "Google Service Account Identifier",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}
