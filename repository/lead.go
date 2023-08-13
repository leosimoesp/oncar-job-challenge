package repository

import (
	"context"

	"github.com/gofrs/uuid"
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
	query := `insert into lead(name, email, phone, vehicle_id) values($1, $2, $3, $4) 
	on conflict(vehicle_id, email) do nothing;
	`
	_, err := lr.db.Exec(ctx, query, lead.Name, lead.Email, lead.Phone, lead.VehicleID)

	return err
}

// FindByVehicle implements domain.LeadRepository.
func (lr leadRepository) FindByVehicle(ctx context.Context, vehicleID uuid.UUID) ([]domain.Lead, error) {
	leads := make([]domain.Lead, 0)

	query := `select l.id, l.name, l.email, l.phone, l.vehicle_id
	from lead l order by l.name
	`
	rows, err := lr.db.Query(ctx, query)

	if err != nil {
		return leads, err
	}
	defer rows.Close()

	for rows.Next() {
		var lead domain.Lead
		err := rows.Scan(&lead.ID, &lead.Name, &lead.Email, &lead.Phone, &lead.VehicleID)
		if err != nil {
			return leads, err
		}
		leads = append(leads, lead)
	}
	return leads, nil
}
