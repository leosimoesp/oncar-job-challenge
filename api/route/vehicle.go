package route

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/leosimoesp/oncar-job-challenge/api/controller"
	"github.com/leosimoesp/oncar-job-challenge/config"
	"github.com/leosimoesp/oncar-job-challenge/repository"
	"github.com/leosimoesp/oncar-job-challenge/usecase"
)

func NewVehicleRouter(cfg config.Config, db *pgxpool.Pool, group *echo.Group) {
	vr := repository.NewVehicleRepository(db)
	vc := &controller.VehicleController{
		VehicleUsecase: usecase.NewVehicleUsecase(cfg, vr),
	}
	group.POST("/vehicles", vc.Find)
}
