package operations

import (
	"context"
	"fmt"
	"time"

	"github.com/elevenode/terraform-provider-expo/internal/client"
	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const pollInterval = 2 * time.Second

func Delete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*client.EASClient)
	id := d.Get("id").(string)

	receipt, err := client.UpdateChannel.Delete(id)
	if err != nil {
		return diag.FromErr(err)
	}

	// Deletion runs as a background job, but Terraform's Delete is synchronous
	// and a dependent branch cannot be removed until the channel is really gone.
	for !receipt.Done() {
		select {
		case <-ctx.Done():
			return diag.FromErr(fmt.Errorf("timed out waiting for channel %s to be deleted; background job %s last reported %s", id, receipt.Id, receipt.State))
		case <-time.After(pollInterval):
		}

		if receipt, err = client.BackgroundJob.GetById(receipt.Id); err != nil {
			return diag.FromErr(err)
		}
	}

	if receipt.State != eas.BackgroundJobStateSuccess {
		message := "no error message was returned"
		if receipt.ErrorMessage != nil {
			message = *receipt.ErrorMessage
		}
		return diag.FromErr(fmt.Errorf("deleting channel %s failed: %s", id, message))
	}

	var diags diag.Diagnostics
	return diags
}
