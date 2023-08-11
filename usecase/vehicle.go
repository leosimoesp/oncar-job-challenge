package usecase

import (
	"context"
	"time"

	"log"

	"github.com/leosimoesp/oncar-job-challenge/config"
	"github.com/leosimoesp/oncar-job-challenge/domain"
)

type vehicleUsecase struct {
	repo domain.VehicleRepository
	cfg  config.Config
}

func NewVehicleUsecase(cfg config.Config, repo domain.VehicleRepository) domain.VehicleUsecase {
	return vehicleUsecase{
		repo: repo,
		cfg:  cfg,
	}
}

// Find implements domain.VehicleUsecase.
func (v vehicleUsecase) Find(ctx context.Context, params domain.FindParams) ([]domain.Vehicle, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Duration(v.cfg.Database.DefaultQueryTimeout)*time.Second)
	defer cancel()
	vehicles, err := v.repo.Find(ctxWithTimeout, params)

	if err != nil {
		log.Default().Printf("Error while finding vehicles: %v", err)
	}

	return vehicles, err
}
