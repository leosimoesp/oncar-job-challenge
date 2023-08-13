package domain

import (
	"fmt"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestLeadRequest_Validate(t *testing.T) {
	shouldReturnNoErrIfValid(t)
	shouldReturnErrIfNameIsEmpty(t)
	shouldReturnErrIfEmailIsEmpty(t)
	shouldReturnErrIfEmailIsInvalid(t)
	shouldReturnErrIfPhoneIsInvalid(t)
}

func shouldReturnErrIfPhoneIsInvalid(t *testing.T) {
	leadReq := LeadRequest{
		Name:      "José",
		Email:     "jose@gmail.com",
		Phone:     "98745-5521",
		VehicleID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
	}
	err := leadReq.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "phone: Invalid phone number. Use (99)99999-9999 or (99)9999-9999.", err.Error())
}

func shouldReturnErrIfEmailIsInvalid(t *testing.T) {
	leadReq := LeadRequest{
		Name:      "José",
		Email:     "jose@y.com",
		Phone:     "(11)98745-5521",
		VehicleID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
	}
	err := leadReq.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "email: must be a valid email address.", err.Error())
}

func shouldReturnErrIfEmailIsEmpty(t *testing.T) {
	leadReq := LeadRequest{
		Name:      "José",
		Email:     "",
		Phone:     "(11)98745-5521",
		VehicleID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
	}
	err := leadReq.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "email: cannot be blank.", err.Error())
}

func shouldReturnNoErrIfValid(t *testing.T) {
	leadReq := LeadRequest{
		Name:      "José",
		Email:     "jose@gmail.com",
		Phone:     "(11)98745-5521",
		VehicleID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
	}
	err := leadReq.Validate()
	assert.Nil(t, err)
	x := uuid.Nil
	fmt.Println(x, leadReq)
}

func shouldReturnErrIfNameIsEmpty(t *testing.T) {
	leadReq := LeadRequest{
		Name:      "",
		Email:     "jose@gmail.com",
		Phone:     "(11)98745-5521",
		VehicleID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
	}
	err := leadReq.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "name: cannot be blank.", err.Error())
}
