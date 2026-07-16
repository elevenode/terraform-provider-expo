package provider

import (
	"context"
	"fmt"
	"github.com/elevenode/terraform-provider-expo/internal/client"
	"github.com/elevenode/terraform-provider-expo/provider/accountvariable"
	androidappcredentials "github.com/elevenode/terraform-provider-expo/provider/android/appcredentials"
	"github.com/elevenode/terraform-provider-expo/provider/android/googleserviceaccountkey"
	"github.com/elevenode/terraform-provider-expo/provider/app"
	"github.com/elevenode/terraform-provider-expo/provider/appvariable"
	iosappcredentials "github.com/elevenode/terraform-provider-expo/provider/ios/appcredentials"
	"github.com/elevenode/terraform-provider-expo/provider/ios/appidentifier"
	"github.com/elevenode/terraform-provider-expo/provider/ios/appstoreapikey"
	"github.com/elevenode/terraform-provider-expo/provider/ios/certificate"
	"github.com/elevenode/terraform-provider-expo/provider/ios/provisioningprofile"
	"github.com/elevenode/terraform-provider-expo/provider/ios/pushkey"
	"github.com/elevenode/terraform-provider-expo/provider/updatebranch"
	"github.com/elevenode/terraform-provider-expo/provider/updatechannel"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Description: "Expo personal access token or robot access token. You can set this via `EXPO TOKEN` environment variable.",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("EXPO_TOKEN", ""),
			},
			"account_name": {
				Description: "Expo user/organization account name. You can set this via `EXPO_ACCOUNT_NAME` environment variable.",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("EXPO_ACCOUNT_NAME", ""),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"expo_app_store_api_key":          appstoreapikey.DataSource(),
			"expo_ios_certificate":            certificate.DataSource(),
			"expo_ios_push_key":               pushkey.DataSource(),
			"expo_google_service_account_key": googleserviceaccountkey.DataSource(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"expo_android_app_credentials":      androidappcredentials.Resource(),
			"expo_app":                          app.Resource(),
			"expo_app_variable":                 appvariable.Resource(),
			"expo_account_variable":             accountvariable.Resource(),
			"expo_ios_app_provisioning_profile": provisioningprofile.Resource(),
			"expo_ios_app_identifier":           appidentifier.Resource(),
			"expo_ios_app_credentials":          iosappcredentials.Resource(),
			"expo_update_branch":                updatebranch.Resource(),
			"expo_update_channel":               updatechannel.Resource(),
		},
		ConfigureContextFunc: func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
			token := d.Get("token").(string)
			accountName := d.Get("account_name").(string)

			var diags diag.Diagnostics

			if token == "" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "token value cannot be an empty string",
					Detail:   "set the token value in the provider configuration or via the EXPO_TOKEN environment variable",
				})
			}

			if accountName == "" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "account_name value cannot be an empty string",
					Detail:   "set the token value in the provider configuration or via the EXPO_ACCOUNT_NAME environment variable",
				})
			}

			if len(diags) > 0 {
				return nil, diags
			}

			client, err := client.NewEASClient(token, accountName)

			if err != nil {
				return nil, diag.Diagnostics{
					diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Failed to initialize EAS client",
						Detail:   fmt.Sprintf("Error: %v", err),
					},
				}
			}

			return client, diags
		},
	}
}
