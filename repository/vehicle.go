package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/leosimoesp/oncar-job-challenge/domain"
	"github.com/leosimoesp/oncar-job-challenge/pkg/input"
)

type vehicleRepository struct {
	db *pgxpool.Pool
}

func NewVehicleRepository(db *pgxpool.Pool) domain.VehicleRepository {
	return &vehicleRepository{
		db: db,
	}
}

func (v vehicleRepository) Find(ctx context.Context, params domain.FindParams) ([]domain.Vehicle, error) {
	vehicles := make([]domain.Vehicle, 0)

	query := `select v.id, v.brand, v.model, v.year, v.price
	from vehicle v order by v.brand, v.model, v.year
	`
	rows, err := v.db.Query(ctx, query)

	if err != nil {
		return vehicles, err
	}
	defer rows.Close()

	for rows.Next() {
		var vehicle domain.Vehicle
		err := rows.Scan(&vehicle.ID, &vehicle.Brand, &vehicle.Model, &vehicle.Year, &vehicle.Price)
		if err != nil {
			return vehicles, err
		}
		vehicle.FmtPrice = input.AddCurrencyMask(vehicle.Price)
		vehicles = append(vehicles, vehicle)
	}
	return vehicles, nil
}
