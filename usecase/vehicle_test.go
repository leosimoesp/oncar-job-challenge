package usecase

import (
	//"context"

	"context"
	"errors"
	"testing"

	"github.com/leosimoesp/oncar-job-challenge/config"
	"github.com/leosimoesp/oncar-job-challenge/domain"
	"github.com/leosimoesp/oncar-job-challenge/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_vehicleUsecase_Find(t *testing.T) {
	succesWhenFindVehicles(t)
	dbErrorFindVehicles(t)
}

func dbErrorFindVehicles(t *testing.T) {
	repo := &mocks.VehicleRepositoryMock{}
	repo.On("Find", mock.Anything, mock.Anything).Return([]domain.Vehicle{}, errors.New("db error"))
	usecase := NewVehicleUsecase(config.Config{}, repo)

	_, err := usecase.Find(context.TODO(), domain.FindParams{})
	assert.NotNil(t, err)
	assert.Equal(t, "db error", err.Error())
}

func succesWhenFindVehicles(t *testing.T) {
	repo := &mocks.VehicleRepositoryMock{}
	repo.On("Find", mock.Anything, mock.Anything).Return([]domain.Vehicle{}, nil)
	usecase := NewVehicleUsecase(config.Config{}, repo)

	_, err := usecase.Find(context.TODO(), domain.FindParams{})
	assert.Nil(t, err)
}
