package route

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/leosimoesp/oncar-job-challenge/api/controller"
	"github.com/leosimoesp/oncar-job-challenge/config"
	"github.com/leosimoesp/oncar-job-challenge/repository"
	"github.com/leosimoesp/oncar-job-challenge/usecase"
)

func NewLeadRouter(cfg config.Config, db *pgxpool.Pool, group *echo.Group) {
	vr := repository.NewLeadRepository(db)
	vc := &controller.LeadController{
		LeadUsecase: usecase.NewLeadUsecase(cfg, vr),
	}
	group.POST("/leads", vc.Save)
	group.GET("/leads/:vehicleId", vc.FindByVehicle)
}
