package apple

import (
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/appbuildcredentials"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/appcredentials"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/appidentifier"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/appstoreapikey"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/certificate"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/provisioningprofile"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/pushkey"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple/team"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/graphql"
)

type Service struct {
	Team                team.Service
	Certificate         certificate.Service
	AppIdentifier       appidentifier.Service
	ProvisioningProfile provisioningprofile.Service
	AppStoreApiKey      appstoreapikey.Service
	AppCredentials      appcredentials.Service
	AppBuildCredentials appbuildcredentials.Service
	PushKey             pushkey.Service
}

func NewService(graphQL graphql.GraphQL) Service {
	return Service{
		Team:                team.NewService(graphQL),
		Certificate:         certificate.NewService(graphQL),
		AppIdentifier:       appidentifier.NewService(graphQL),
		ProvisioningProfile: provisioningprofile.NewService(graphQL),
		AppStoreApiKey:      appstoreapikey.NewService(graphQL),
		AppCredentials:      appcredentials.NewService(graphQL),
		AppBuildCredentials: appbuildcredentials.NewService(graphQL),
		PushKey:             pushkey.NewService(graphQL),
	}
}
