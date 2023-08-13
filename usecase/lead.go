package usecase

import (
	"context"
	"log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/leosimoesp/oncar-job-challenge/config"
	"github.com/leosimoesp/oncar-job-challenge/domain"
)

type leadUsecase struct {
	repo domain.LeadRepository
	cfg  config.Config
}

func NewLeadUsecase(cfg config.Config, repo domain.LeadRepository) domain.LeadUsecase {
	return leadUsecase{
		repo: repo,
		cfg:  cfg,
	}
}

// Save implements domain.LeadUsecase.
func (l leadUsecase) Save(ctx context.Context, leadReq domain.LeadRequest) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Duration(l.cfg.Database.DefaultQueryTimeout)*time.Second)
	defer cancel()
	err := l.repo.Save(ctxWithTimeout, leadReq.ToLead())

	if err != nil {
		log.Default().Printf("Error while finding vehicles: %v", err)
	}
	return err
}

// FindByVehicle implements domain.LeadUsecase.
func (l leadUsecase) FindByVehicle(ctx context.Context, vehicleID string) ([]domain.Lead, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Duration(l.cfg.Database.DefaultQueryTimeout)*time.Second)
	defer cancel()
	leads, err := l.repo.FindByVehicle(ctxWithTimeout, uuid.FromStringOrNil(vehicleID))

	if err != nil {
		log.Default().Printf("Error while finding leads from vehicle: %v", err)
	}
	return leads, err
}
