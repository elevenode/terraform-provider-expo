package android

import (
	"terraform-provider-eas/internal/eas/internal/api/android/appbuildcredentials"
	"terraform-provider-eas/internal/eas/internal/api/android/appcredentials"
	"terraform-provider-eas/internal/eas/internal/api/android/fcmkey"
	"terraform-provider-eas/internal/eas/internal/api/android/googleserviceaccountkey"
	"terraform-provider-eas/internal/eas/internal/api/android/keystore"
	"terraform-provider-eas/internal/eas/internal/graphql"
)

type Service struct {
	GoogleServiceAccountKey googleserviceaccountkey.Service
	AppCredentials          appcredentials.Service
	AppBuildCredentials     appbuildcredentials.Service
	FCMKey                  fcmkey.Service
	Keystore                keystore.Service
}

func NewService(graphQL graphql.GraphQL) Service {
	return Service{
		GoogleServiceAccountKey: googleserviceaccountkey.NewService(graphQL),
		AppCredentials:          appcredentials.NewService(graphQL),
		AppBuildCredentials:     appbuildcredentials.NewService(graphQL),
		FCMKey:                  fcmkey.NewService(graphQL),
		Keystore:                keystore.NewService(graphQL),
	}
}
