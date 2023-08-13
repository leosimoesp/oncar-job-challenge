package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leosimoesp/oncar-job-challenge/domain"
	"github.com/leosimoesp/oncar-job-challenge/pkg/validator"
)

type LeadController struct {
	LeadUsecase domain.LeadUsecase
}

func (lc LeadController) Save(c echo.Context) error {
	ctx := c.Request().Context()
	var leadReq domain.LeadRequest

	if err := c.Bind(&leadReq); err != nil {
		c.Error(err)
		return err
	}
	if e := leadReq.Validate(); e != nil {
		return c.JSON(http.StatusUnprocessableEntity, e)
	}
	err := lc.LeadUsecase.Save(ctx, leadReq)
	if err != nil {
		c.Error(err)
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (lc LeadController) FindByVehicle(c echo.Context) error {
	ctx := c.Request().Context()
	vehicleID := c.Param("vehicleId")

	if err := validator.ValidateParam(vehicleID, "vehicleId not found in path"); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	leads, err := lc.LeadUsecase.FindByVehicle(ctx, vehicleID)
	if err != nil {
		c.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, leads)
}
