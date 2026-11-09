package operations

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Read keeps the stored state as is: the EAS API has no query to fetch an
// Android keystore by id, only to create and delete one.
func Read(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	return nil
}
