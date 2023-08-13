package domain

import (
	"context"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofrs/uuid"
	"github.com/leosimoesp/oncar-job-challenge/pkg/input"
	"github.com/leosimoesp/oncar-job-challenge/pkg/validator"
)

type Lead struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	VehicleID uuid.UUID `json:"vehicleId"`
	ID        uuid.UUID `json:"id"`
}

type LeadRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	VehicleID string `json:"vehicleId"`
}

func (lr LeadRequest) Validate() error {
	lr.clear()
	return validation.ValidateStruct(&lr,
		validation.Field(&lr.Name, validation.Required, validation.Length(5, 150)),
		validation.Field(&lr.Phone, validation.Required, is.Digit),
		validation.Field(&lr.Phone, validation.By(validator.ValidatePhone)),
		validation.Field(&lr.Email, validation.Required, is.Email),
		validation.Field(&lr.VehicleID, validation.Required.Error("VehicleID is mandatory"), is.UUID),
	)
}
func (lr *LeadRequest) ToLead() Lead {
	lr.clear()
	return Lead{
		Name:      lr.Name,
		Email:     lr.Email,
		Phone:     lr.Phone,
		VehicleID: uuid.FromStringOrNil(lr.VehicleID),
	}
}

func (lr *LeadRequest) clear() {
	name := strings.ReplaceAll(lr.Name, " ", "")
	email := strings.ReplaceAll(lr.Email, " ", "")
	phone := input.RemoveWithPattern(lr.Phone, "[^0-9]+")
	lr.Email = email
	lr.Phone = phone
	lr.Name = name
}

type LeadRepository interface {
	Save(ctx context.Context, lead Lead) error
}

type LeadUsecase interface {
	Save(ctx context.Context, lead LeadRequest) error
}
