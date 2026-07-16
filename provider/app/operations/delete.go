package operations

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/elevenode/terraform-provider-expo/internal/client"
	"github.com/elevenode/terraform-provider-expo/internal/eas"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	deletionPollDelay    = 2 * time.Second
	deletionPollInterval = 5 * time.Second
)

// EAS deletes apps asynchronously, so scheduling only hands back a background job receipt.
// Terraform's Delete contract is synchronous, so poll the receipt to a terminal state.
func Delete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*client.EASClient)
	id := d.Id()

	receipt, err := client.App.Delete(id)
	if err != nil {
		return diag.FromErr(scheduleError(client, id, err))
	}
	if receipt == nil {
		return diag.Errorf("EAS returned no background job receipt when scheduling the deletion of app %s. Check %s to see whether it was deleted", id, dashboardUrl(client))
	}

	if err := waitForDeletion(ctx, client, id, receipt.Id, d.Timeout(schema.TimeoutDelete)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func waitForDeletion(ctx context.Context, client *client.EASClient, appId string, receiptId string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(eas.BackgroundJobStateQueued),
			string(eas.BackgroundJobStateInProgress),
			// A FAILURE that EAS will retry has not settled yet, so keep waiting on it.
			string(eas.BackgroundJobStateFailure),
		},
		Target:     []string{string(eas.BackgroundJobStateSuccess)},
		Timeout:    timeout,
		Delay:      deletionPollDelay,
		MinTimeout: deletionPollInterval,
		Refresh: func() (any, string, error) {
			receipt, err := client.BackgroundJob.Get(receiptId)
			if err != nil {
				return nil, "", fmt.Errorf("failed to read the status of deletion job %s for app %s: %w", receiptId, appId, err)
			}
			if receipt == nil {
				return nil, "", fmt.Errorf("deletion job %s for app %s was not found", receiptId, appId)
			}
			if receipt.State == eas.BackgroundJobStateFailure && !receipt.WillRetry {
				return nil, "", failureError(client, appId, receipt)
			}
			return receipt, string(receipt.State), nil
		},
	}

	if _, err := stateConf.WaitForStateContext(ctx); err != nil {
		var timeoutErr *retry.TimeoutError
		if errors.As(err, &timeoutErr) {
			return fmt.Errorf(
				"timed out after %s waiting for EAS to delete app %s. The deletion job may still finish on the Expo side, check %s before retrying",
				timeout, appId, dashboardUrl(client),
			)
		}
		return err
	}

	return nil
}

func dashboardUrl(client *client.EASClient) string {
	return fmt.Sprintf("https://expo.dev/accounts/%s/projects", client.AccountName)
}

// Deleting an app requires elevated privileges, by far the likeliest reason for a rejection.
func scheduleError(client *client.EASClient, appId string, err error) error {
	return fmt.Errorf(
		"failed to schedule the deletion of app %s: %w\n\nDeleting an app requires elevated privileges on the Expo account. Check that the token the provider uses belongs to an owner or admin of account %q, or delete the app manually at %s and drop it from the state with `terraform state rm`",
		appId, err, client.AccountName, dashboardUrl(client),
	)
}

func failureError(client *client.EASClient, appId string, receipt *eas.BackgroundJobData) error {
	detail := receipt.ErrorMessage
	if detail == "" {
		detail = "EAS returned no error message"
	}
	if receipt.ErrorCode != "" {
		detail = fmt.Sprintf("%s (%s)", detail, receipt.ErrorCode)
	}

	return fmt.Errorf(
		"EAS failed to delete app %s after %d attempt(s): %s. The app is kept in the Terraform state, check %s and retry",
		appId, receipt.Tries, detail, dashboardUrl(client),
	)
}
