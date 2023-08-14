package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leosimoesp/oncar-job-challenge/domain"
)

type VehicleController struct {
	VehicleUsecase domain.VehicleUsecase
}

func (vc VehicleController) Find(c echo.Context) error {
	ctx := c.Request().Context()

	var params domain.FindParams

	if err := c.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	vehicles, err := vc.VehicleUsecase.Find(ctx, params)

	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, vehicles)
}
