package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/leosimoesp/oncar-job-challenge/domain"
	"github.com/leosimoesp/oncar-job-challenge/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLeadController_Save(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rec := httptest.NewRecorder()

		var expectedResult error

		leadUsecase := mocks.LeadUsecaseMock{}
		leadUsecase.On("Save", mock.Anything, mock.Anything).Return(expectedResult)

		lc := LeadController{
			LeadUsecase: &leadUsecase,
		}

		e := echo.New()
		leadReq := domain.LeadRequest{
			Name:      "José",
			Email:     "jose@gmail.com",
			Phone:     "(11)98745-5521",
			VehicleID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		}
		body, err := json.Marshal(leadReq)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/leads", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e.NewContext(req, rec)

		grApi := e.Group("/api")
		grApi.POST("/leads", lc.Save)
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	})

	t.Run("bind_error", func(t *testing.T) {
		rec := httptest.NewRecorder()

		leadUsecase := mocks.LeadUsecaseMock{}

		lc := LeadController{
			LeadUsecase: &leadUsecase,
		}

		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/api/leads", bytes.NewReader([]byte(`{`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e.NewContext(req, rec)

		grApi := e.Group("/api")
		grApi.POST("/leads", lc.Save)
		e.ServeHTTP(rec, req)

		jsonMsgErr, _ := json.Marshal(domain.MsgErr{Message: "unexpected EOF"})
		expectedErr := errors.New(string(jsonMsgErr))

		err := errors.New(rec.Body.String())
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, expectedErr.Error()+"\n", err.Error())
	})

	t.Run("invalid_request_error", func(t *testing.T) {
		rec := httptest.NewRecorder()

		var expectedResult error

		leadUsecase := mocks.LeadUsecaseMock{}
		leadUsecase.On("Save", mock.Anything, mock.Anything).Return(expectedResult)

		lc := LeadController{
			LeadUsecase: &leadUsecase,
		}

		e := echo.New()
		leadReq := domain.LeadRequest{
			Name:      "José",
			Email:     "",
			Phone:     "(11)98745-5521",
			VehicleID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		}
		body, err := json.Marshal(leadReq)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/leads", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e.NewContext(req, rec)

		grApi := e.Group("/api")
		grApi.POST("/leads", lc.Save)
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Equal(t, "{\"email\":\"cannot be blank\"}\n", rec.Body.String())
	})

	t.Run("err_when_call_usecase", func(t *testing.T) {
		rec := httptest.NewRecorder()

		leadUsecase := mocks.LeadUsecaseMock{}
		leadUsecase.On("Save", mock.Anything, mock.Anything).Return(errors.New("db error"))

		lc := LeadController{
			LeadUsecase: &leadUsecase,
		}

		e := echo.New()
		leadReq := domain.LeadRequest{
			Name:      "José",
			Email:     "jose@gmail.com",
			Phone:     "(11)98745-5521",
			VehicleID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		}
		body, err := json.Marshal(leadReq)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/leads", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e.NewContext(req, rec)

		grApi := e.Group("/api")
		grApi.POST("/leads", lc.Save)
		e.ServeHTTP(rec, req)

		jsonMsgErr, _ := json.Marshal(domain.MsgErr{Message: "Internal Server Error"})
		expectedErr := errors.New(string(jsonMsgErr))
		err = errors.New(rec.Body.String())

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, expectedErr.Error()+"\n", err.Error())
	})
}

func TestLeadController_FindByVehicle(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rec := httptest.NewRecorder()

		expectedResult := []domain.Lead{
			{
				Name:      "José",
				Email:     "jse@gmail.com",
				Phone:     "11977415474",
				VehicleID: uuid.FromStringOrNil("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
			},
		}

		leadUsecase := mocks.LeadUsecaseMock{}
		leadUsecase.On("FindByVehicle", mock.Anything, mock.Anything).Return(expectedResult, nil)

		lc := LeadController{
			LeadUsecase: &leadUsecase,
		}

		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/api/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		echoCtx := e.NewContext(req, rec)
		echoCtx.SetPath("leads/:vehicleId")
		echoCtx.SetParamNames("vehicleId")
		echoCtx.SetParamValues("6ba7b810-9dad-11d1-80b4-00c04fd430c8")

		grApi := e.Group("/api")
		grApi.GET("/leads", lc.FindByVehicle)

		err := lc.FindByVehicle(echoCtx)
		assert.Nil(t, err)

		expectedResultAsJson, _ := json.Marshal(expectedResult)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResultAsJson)+"\n", rec.Body.String())
	})

	t.Run("invalid_request_error", func(t *testing.T) {
		rec := httptest.NewRecorder()

		expectedResult := []domain.Lead{
			{
				Name:      "José",
				Email:     "jse@gmail.com",
				Phone:     "11977415474",
				VehicleID: uuid.FromStringOrNil("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
			},
		}

		leadUsecase := mocks.LeadUsecaseMock{}
		leadUsecase.On("FindByVehicle", mock.Anything, mock.Anything).Return(expectedResult, nil)

		lc := LeadController{
			LeadUsecase: &leadUsecase,
		}

		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/api/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		echoCtx := e.NewContext(req, rec)
		echoCtx.SetPath("leads/:vehicleId")

		grApi := e.Group("/api")
		grApi.GET("/leads", lc.FindByVehicle)

		err := lc.FindByVehicle(echoCtx)
		assert.NotNil(t, err)
		assert.Equal(t, "code=400, message=vehicleId not found in path", err.Error())
	})

	t.Run("err_when_call_usecase", func(t *testing.T) {
		rec := httptest.NewRecorder()

		expectedResult := []domain.Lead{}

		leadUsecase := mocks.LeadUsecaseMock{}
		leadUsecase.On("FindByVehicle", mock.Anything, mock.Anything).Return(expectedResult, errors.New("db timeout"))

		lc := LeadController{
			LeadUsecase: &leadUsecase,
		}

		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/api/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		echoCtx := e.NewContext(req, rec)
		echoCtx.SetPath("leads/:vehicleId")
		echoCtx.SetParamNames("vehicleId")
		echoCtx.SetParamValues("6ba7b810-9dad-11d1-80b4-00c04fd430c8")

		grApi := e.Group("/api")
		grApi.GET("/leads", lc.FindByVehicle)

		err := lc.FindByVehicle(echoCtx)
		assert.NotNil(t, err)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "{\"message\":\"Internal Server Error\"}\n", rec.Body.String())
	})
}
