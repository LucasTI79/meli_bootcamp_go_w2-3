package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type LocalityServiceMock struct {
	mock.Mock
}

type LocalityRepositoryMock struct {
	mock.Mock
}

func (m *LocalityRepositoryMock) ExistsById(ctx context.Context, id int) bool {
	args := m.Called(ctx, id)
	return args.Get(0).(bool)
}