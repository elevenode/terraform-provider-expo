package machinebox

import (
	"context"

	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/mock"
)

type MachineBoxGraphQLClientMock struct {
	mock.Mock
}

func (m *MachineBoxGraphQLClientMock) Run(ctx context.Context, req *graphql.Request, resp any) error {
	args := m.Called(ctx, req, resp)
	return args.Error(0)
}

var _ MachineboxGraphqlClient = (*MachineBoxGraphQLClientMock)(nil)
