package operations

import (
	"context"
	"fmt"
	"strings"

	"github.com/elevenode/terraform-provider-expo/internal/client"
	"github.com/elevenode/terraform-provider-expo/internal/eas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Import takes an "<app_id>/<channel_name>" address, since an update channel is
// only addressable by name within an app.
func Import(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	client := m.(*client.EASClient)

	appId, name, found := strings.Cut(d.Id(), "/")
	if !found || appId == "" || name == "" {
		return nil, fmt.Errorf("invalid import id %q, expected \"<app_id>/<channel_name>\"", d.Id())
	}

	data, err := client.UpdateChannel.GetByName(eas.GetByNameUpdateChannelData{
		AppId: appId,
		Name:  name,
	})
	if err != nil {
		return nil, err
	}

	d.SetId(data.Id)

	if err := d.Set("id", data.Id); err != nil {
		return nil, err
	}
	if err := d.Set("name", data.Name); err != nil {
		return nil, err
	}
	if err := d.Set("branch_mapping", data.BranchMapping); err != nil {
		return nil, err
	}
	if err := d.Set("app_id", appId); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
