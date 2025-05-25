package eas

import (
	"github.com/fintreal/eas-sdk-go/internal/api/account"
	"github.com/fintreal/eas-sdk-go/internal/api/android"
	"github.com/fintreal/eas-sdk-go/internal/api/app"
	"github.com/fintreal/eas-sdk-go/internal/api/apple"
	"github.com/fintreal/eas-sdk-go/internal/api/appvariable"
	"github.com/fintreal/eas-sdk-go/internal/api/me"
	"github.com/fintreal/eas-sdk-go/internal/graphql"
)

// EASClient capable of interacting with Expo EAS GraphQL API
type EASClient struct {
	Me          me.Service
	App         app.Service
	AppVariable appvariable.Service
	Account     account.AccountService
	Apple       apple.Service
	Android     android.Service
}

// EASClient capable of interacting with Expo EAS GraphQL API
//
// @token Expo Personal Access Token or Robot Access Token
func NewEASClient(token string) *EASClient {
	if token == "" {
		panic("expo token can't be an empty string")
	}
	graphql := graphql.NewGraphQL(token)
	return &EASClient{
		Me:          me.NewService(graphql),
		App:         app.NewService(graphql),
		AppVariable: appvariable.NewService(graphql),
		Account:     account.NewAccountService(graphql),
		Apple:       apple.NewService(graphql),
		Android:     android.NewService(graphql),
	}
}
