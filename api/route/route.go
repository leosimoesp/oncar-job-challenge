package route

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/leosimoesp/oncar-job-challenge/config"
)

func SetupApiRoutes(cfg config.Config, db *pgxpool.Pool, e *echo.Echo) {
	apiGroup := e.Group("/api")
	NewVehicleRouter(cfg, db, apiGroup)
	NewLeadRouter(cfg, db, apiGroup)
}
