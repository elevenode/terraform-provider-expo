package apple

import (
	"terraform-provider-eas/internal/eas/internal/api/apple/appbuildcredentials"
	"terraform-provider-eas/internal/eas/internal/api/apple/appcredentials"
	"terraform-provider-eas/internal/eas/internal/api/apple/appidentifier"
	"terraform-provider-eas/internal/eas/internal/api/apple/appstoreapikey"
	"terraform-provider-eas/internal/eas/internal/api/apple/certificate"
	"terraform-provider-eas/internal/eas/internal/api/apple/provisioningprofile"
	"terraform-provider-eas/internal/eas/internal/api/apple/pushkey"
	"terraform-provider-eas/internal/eas/internal/api/apple/team"
	"terraform-provider-eas/internal/eas/internal/graphql"
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
