package machinebox

import (
	"context"

	"github.com/machinebox/graphql"
)

type MachineboxGraphqlClient interface {
	Run(ctx context.Context, req *graphql.Request, res any) error
}

var _ MachineboxGraphqlClient = (*graphql.Client)(nil)
