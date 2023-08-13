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
