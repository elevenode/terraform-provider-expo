package eas

import (
	"terraform-provider-eas/internal/eas/internal/api/account"
	"terraform-provider-eas/internal/eas/internal/api/accountvariable"
	"terraform-provider-eas/internal/eas/internal/api/android"
	"terraform-provider-eas/internal/eas/internal/api/app"
	"terraform-provider-eas/internal/eas/internal/api/apple"
	"terraform-provider-eas/internal/eas/internal/api/appvariable"
	"terraform-provider-eas/internal/eas/internal/api/me"
	"terraform-provider-eas/internal/eas/internal/graphql"
)

// EASClient capable of interacting with Expo EAS GraphQL API
type EASClient struct {
	Me              me.Service
	App             app.Service
	AppVariable     appvariable.Service
	Account         account.AccountService
	AccountVariable accountvariable.Service
	Apple           apple.Service
	Android         android.Service
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
		Me:              me.NewService(graphql),
		App:             app.NewService(graphql),
		AppVariable:     appvariable.NewService(graphql),
		Account:         account.NewAccountService(graphql),
		AccountVariable: accountvariable.NewService(graphql),
		Apple:           apple.NewService(graphql),
		Android:         android.NewService(graphql),
	}
}
