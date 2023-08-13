package mocks

import (
	"context"

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
