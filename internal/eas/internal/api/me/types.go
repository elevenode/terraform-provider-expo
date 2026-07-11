package me

import (
	"terraform-provider-eas/internal/eas/internal/graphql"
)

type Data struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type Service interface {
	Get() (*Data, error)
}

type service struct {
	graphql graphql.GraphQL
}

var _ Service = (*service)(nil)

func NewService(graphql graphql.GraphQL) Service {
	return &service{
		graphql: graphql,
	}
}
