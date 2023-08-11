package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leosimoesp/oncar-job-challenge/config"
)

func SetupApiRoutes(cfg config.Config, e *echo.Echo) {
	apiGroup := e.Group("/api")
	apiGroup.GET("/test", test)
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "Test Setup")
}
