package apple

import (
	"github.com/fintreal/eas-sdk-go/internal/api/apple/appbuildcredentials"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/appcredentials"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/appidentifier"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/appstoreapikey"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/certificate"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/provisioningprofile"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/pushkey"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/team"
	"github.com/fintreal/eas-sdk-go/internal/graphql"
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
