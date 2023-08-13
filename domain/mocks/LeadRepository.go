package mocks

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/leosimoesp/oncar-job-challenge/domain"
	"github.com/stretchr/testify/mock"
)

type LeadRepositoryMock struct {
	mock.Mock
}

func (m *LeadRepositoryMock) Save(ctx context.Context, lead domain.Lead) error {
	args := m.Called(ctx, lead)
	return args.Error(0)
}

func (m *LeadRepositoryMock) FindByVehicle(ctx context.Context, vehicleID uuid.UUID) ([]domain.Lead, error) {
	args := m.Called(ctx, vehicleID)
	return args.Get(0).([]domain.Lead), args.Error(1)
}
