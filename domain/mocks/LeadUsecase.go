package mocks

import (
	"context"

	"github.com/leosimoesp/oncar-job-challenge/domain"
	"github.com/stretchr/testify/mock"
)

type LeadUsecaseMock struct {
	mock.Mock
}

func (m *LeadUsecaseMock) Save(ctx context.Context, lead domain.LeadRequest) error {
	args := m.Called(ctx, lead)
	return args.Error(0)
}

func (m *LeadUsecaseMock) FindByVehicle(ctx context.Context, vehicleID string) ([]domain.Lead, error) {
	args := m.Called(ctx, vehicleID)
	return args.Get(0).([]domain.Lead), args.Error(1)
}
