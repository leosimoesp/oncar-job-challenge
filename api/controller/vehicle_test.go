package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/leosimoesp/oncar-job-challenge/domain"
	"github.com/leosimoesp/oncar-job-challenge/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestVehicleController_Find(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rec := httptest.NewRecorder()

		expectedResult := []domain.Vehicle{{
			Brand: "Fiat",
			Model: "ZasX32",
			Year:  2019,
			Price: 100000,
		}}

		vehicleUsecase := mocks.VehicleUsecaseMock{}
		vehicleUsecase.On("Find", mock.Anything, mock.Anything).Return(expectedResult, nil)

		vc := VehicleController{
			VehicleUsecase: &vehicleUsecase,
		}

		e := echo.New()
		var params domain.FindParams
		body, err := json.Marshal(params)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/vehicles", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e.NewContext(req, rec)

		grApi := e.Group("/api")
		grApi.POST("/vehicles", vc.Find)
		e.ServeHTTP(rec, req)

		expectedResultAsJson, _ := json.Marshal(expectedResult)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResultAsJson)+"\n", rec.Body.String())

	})

	t.Run("bind_error", func(t *testing.T) {
		rec := httptest.NewRecorder()

		vehicleUsecase := mocks.VehicleUsecaseMock{}

		vc := VehicleController{
			VehicleUsecase: &vehicleUsecase,
		}

		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/api/vehicles", bytes.NewReader([]byte(`{`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e.NewContext(req, rec)

		grApi := e.Group("/api")
		grApi.POST("/vehicles", vc.Find)
		e.ServeHTTP(rec, req)

		jsonMsgErr, _ := json.Marshal(domain.MsgErr{Message: "code=400, message=unexpected EOF, internal=unexpected EOF"})
		expectedErr := errors.New(string(jsonMsgErr))

		err := errors.New(rec.Body.String())
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, expectedErr.Error()+"\n", err.Error())
	})

	t.Run("err_when_call_usecase", func(t *testing.T) {
		rec := httptest.NewRecorder()

		expectedResult := []domain.Vehicle{}

		vehicleUsecase := mocks.VehicleUsecaseMock{}
		vehicleUsecase.On("Find", mock.Anything, mock.Anything).Return(expectedResult, errors.New("db error"))

		vc := VehicleController{
			VehicleUsecase: &vehicleUsecase,
		}

		e := echo.New()
		var params domain.FindParams
		body, err := json.Marshal(params)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/vehicles", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e.NewContext(req, rec)

		grApi := e.Group("/api")
		grApi.POST("/vehicles", vc.Find)
		e.ServeHTTP(rec, req)

		jsonMsgErr, _ := json.Marshal(domain.MsgErr{Message: "Internal Server Error"})
		expectedErr := errors.New(string(jsonMsgErr))
		err = errors.New(rec.Body.String())

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, expectedErr.Error()+"\n", err.Error())
	})
}

type EchoContextMock struct {
	mock.Mock
}

func (m *EchoContextMock) JSON(code int, i interface{}) error {
	args := m.Called(code, i)
	return args.Error(0)
}
