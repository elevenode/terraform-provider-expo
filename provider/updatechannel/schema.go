package updatechannel

import (
	"github.com/elevenode/terraform-provider-expo/provider/updatechannel/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   operations.Read,
		CreateContext: operations.Create,
		UpdateContext: operations.Update,
		DeleteContext: operations.Delete,
		Importer:      &schema.ResourceImporter{StateContext: operations.Import},
		Schema: map[string]*schema.Schema{
			"app_id": {
				Description: "The id of the app the update channel belongs to",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "The name of the update channel. This is the value referenced by a build profile's `channel` in eas.json.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"id": {
				Description: "The id of the update channel",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"branch_mapping": {
				Description:      "Stringified JSON describing which branches the channel routes updates to. Use `jsonencode()`. See https://docs.expo.dev/eas-update/channel-surfing/ for the mapping format, including rollouts.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},
		},
	}
}
