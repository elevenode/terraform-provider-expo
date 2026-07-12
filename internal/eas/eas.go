package eas

import (
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/account"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/accountvariable"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/android"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/app"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/apple"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/appvariable"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/api/me"
	"github.com/elevenode/terraform-provider-expo/internal/eas/internal/graphql"
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
