package operations

import (
	"fmt"

	"github.com/elevenode/terraform-provider-expo/internal/client"
	"github.com/elevenode/terraform-provider-expo/internal/eas"
)

// createBuildCredentials creates each configured build credential under the
// given app credentials and returns them in the shape stored in state.
func createBuildCredentials(client *client.EASClient, appCredentialsId string, configured []any) ([]any, error) {
	created := make([]any, 0, len(configured))

	for _, raw := range configured {
		buildCredential, ok := raw.(map[string]any)
		if !ok {
			return created, fmt.Errorf("error while parsing build credentials config")
		}

		data, err := client.Android.AppBuildCredentials.Create(eas.CreateAndroidAppBuildCredentialsData{
			AppCredentialsId: appCredentialsId,
			Name:             buildCredential["name"].(string),
			KeystoreId:       buildCredential["keystore_id"].(string),
		})
		if err != nil {
			return created, err
		}

		created = append(created, map[string]any{
			"id":          data.Id,
			"name":        data.Name,
			"keystore_id": data.KeystoreId,
		})
	}

	return created, nil
}
