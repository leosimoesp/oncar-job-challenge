package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/leosimoesp/oncar-job-challenge/domain"
)

type leadRepository struct {
	db *pgxpool.Pool
}

func NewLeadRepository(db *pgxpool.Pool) domain.LeadRepository {
	return &leadRepository{
		db: db,
	}
}

// Save implements domain.LeadRepository.
func (lr leadRepository) Save(ctx context.Context, lead domain.Lead) error {

	query := `insert into lead(name, email, phone, vehicle_id) values($1, $2, $3, $4)`
	_, err := lr.db.Exec(ctx, query, lead.Name, lead.Email, lead.Phone, lead.VehicleID)

	return err
}
