package appvariable

import (
	"context"
	"fmt"
	"strings"

	"github.com/elevenode/terraform-provider-expo/provider/appvariable/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Description:   "An EAS environment variable scoped to an app. Set `type = \"FILE_BASE64\"` to create a file variable.",
		CustomizeDiff: validateFileName,
		ReadContext:   operations.Read,
		CreateContext: operations.Create,
		UpdateContext: operations.Update,
		DeleteContext: operations.Delete,
		Schema: map[string]*schema.Schema{
			"app_id": {
				Description: "The id of the app for the environment variable",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "The name of the app for the environment variable",
				Type:        schema.TypeString,
				Required:    true,
			},
			"id": {
				Description: "The id of the environment variable",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"value": {
				Description: "The value of the environment variable. For a `FILE_BASE64` variable, the base64-encoded file content, e.g. `filebase64(\"google-services.json\")`.",
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
			},
			"type": {
				Description:  "The type of the environment variable: `STRING` (default) or `FILE_BASE64` for a file variable.",
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "STRING",
				ValidateFunc: validation.StringInSlice([]string{"STRING", "FILE_BASE64"}, false),
			},
			"file_name": {
				Description: "The file name of a `FILE_BASE64` variable, e.g. `google-services.json`.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"visibility": {
				Description:  "The visibility of the app for the environment variable",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"PUBLIC", "SENSITIVE", "SECRET"}, false),
			},
			"environments": {
				Description: "The environments of the app for the environment variable",
				Type:        schema.TypeSet,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"development", "preview", "production"}, true),
					StateFunc: func(val any) string {
						return strings.ToLower(val.(string))
					},
				},
				Required: true,
			},
		},
	}
}

// validateFileName rejects file_name on STRING variables, where EAS ignores it.
func validateFileName(ctx context.Context, d *schema.ResourceDiff, m any) error {
	if d.Get("type").(string) != "FILE_BASE64" && d.HasChange("file_name") && d.Get("file_name").(string) != "" {
		return fmt.Errorf("file_name can only be set when type is FILE_BASE64")
	}
	return nil
}
