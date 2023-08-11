package domain

import (
	"context"

	"github.com/gofrs/uuid"
)

// Price in cents
type Vehicle struct {
	Brand string    `json:"brand"`
	Model string    `json:"model"`
	Year  int       `json:"year"`
	Price int64     `json:"price"`
	ID    uuid.UUID `json:"id"`
}

type FindParams struct {
	ID     uuid.UUID `json:"id"`
	Cursor uuid.UUID `json:"cursor"`
}

type VehicleRepository interface {
	Find(ctx context.Context, params FindParams) ([]Vehicle, error)
}

type VehicleUsecase interface {
	Find(ctx context.Context, params FindParams) ([]Vehicle, error)
}
