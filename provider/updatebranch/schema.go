package updatebranch

import (
	"github.com/elevenode/terraform-provider-expo/provider/updatebranch/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   operations.Read,
		CreateContext: operations.Create,
		DeleteContext: operations.Delete,
		Importer:      &schema.ResourceImporter{StateContext: operations.Import},
		Schema: map[string]*schema.Schema{
			"app_id": {
				Description: "The id of the app the update branch belongs to",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "The name of the update branch",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"id": {
				Description: "The id of the update branch",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}
