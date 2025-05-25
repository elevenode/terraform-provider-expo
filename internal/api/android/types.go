package android

import (
	"github.com/fintreal/eas-sdk-go/internal/api/android/appbuildcredentials"
	"github.com/fintreal/eas-sdk-go/internal/api/android/appcredentials"
	"github.com/fintreal/eas-sdk-go/internal/api/android/fcmkey"
	"github.com/fintreal/eas-sdk-go/internal/api/android/googleserviceaccountkey"
	"github.com/fintreal/eas-sdk-go/internal/api/android/keystore"
	"github.com/fintreal/eas-sdk-go/internal/graphql"
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
