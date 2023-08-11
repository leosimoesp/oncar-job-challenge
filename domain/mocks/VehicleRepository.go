package mocks

import (
	"context"

	"github.com/leosimoesp/oncar-job-challenge/domain"
	"github.com/stretchr/testify/mock"
)

type VehicleRepositoryMock struct {
	mock.Mock
}

func (m *VehicleRepositoryMock) Find(ctx context.Context, params domain.FindParams) ([]domain.Vehicle, error) {
	args := m.Called(ctx, params)
	return args.Get(0).([]domain.Vehicle), args.Error(1)
}
