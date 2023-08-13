package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/leosimoesp/oncar-job-challenge/config"
	"github.com/leosimoesp/oncar-job-challenge/domain"
	"github.com/leosimoesp/oncar-job-challenge/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_leadUsecase_Save(t *testing.T) {
	succesWhenSaveLead(t)
	dbErrorSaveLead(t)
}

func dbErrorSaveLead(t *testing.T) {
	repo := &mocks.LeadRepositoryMock{}
	repo.On("Save", mock.Anything, mock.Anything).Return(errors.New("db error"))
	usecase := NewLeadUsecase(config.Config{}, repo)

	err := usecase.Save(context.TODO(), domain.LeadRequest{
		Name:      "Name",
		Email:     "test@gmail.com",
		Phone:     "11784577741",
		VehicleID: "123e4567-e89b-12d3-a456-426655440000",
	})
	assert.NotNil(t, err)
	assert.Equal(t, "db error", err.Error())
}

func succesWhenSaveLead(t *testing.T) {
	repo := &mocks.LeadRepositoryMock{}
	repo.On("Save", mock.Anything, mock.Anything).Return(nil)
	usecase := NewLeadUsecase(config.Config{}, repo)

	err := usecase.Save(context.TODO(), domain.LeadRequest{
		Name:      "Name",
		Email:     "test@gmail.com",
		Phone:     "11784577741",
		VehicleID: "123e4567-e89b-12d3-a456-426655440000",
	})
	assert.Nil(t, err)
}

func Test_leadUsecase_FindByVehicle(t *testing.T) {
	succesWhenFindLeadsVehicle(t)
	dbErrorFindLeadsVehicle(t)
}

func dbErrorFindLeadsVehicle(t *testing.T) {
	repo := &mocks.LeadRepositoryMock{}
	repo.On("FindByVehicle", mock.Anything, mock.Anything).Return([]domain.Lead{}, errors.New("db error"))
	usecase := NewLeadUsecase(config.Config{}, repo)

	_, err := usecase.FindByVehicle(context.TODO(), "6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	assert.NotNil(t, err)
	assert.Equal(t, "db error", err.Error())
}

func succesWhenFindLeadsVehicle(t *testing.T) {
	repo := &mocks.LeadRepositoryMock{}
	repo.On("FindByVehicle", mock.Anything, mock.Anything).Return([]domain.Lead{}, nil)
	usecase := NewLeadUsecase(config.Config{}, repo)

	_, err := usecase.FindByVehicle(context.TODO(), "6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	assert.Nil(t, err)
}
