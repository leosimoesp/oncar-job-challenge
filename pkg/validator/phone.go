package validator

import (
	"fmt"

	"github.com/leosimoesp/oncar-job-challenge/pkg/input"
)

const (
	PhoneMinLength = 10
	PhoneMaxLength = 11
)

func ValidatePhone(value interface{}) error {
	rawValue, _ := value.(string)
	phoneOnlyNumbers := input.RemoveWithPattern(rawValue, "[^0-9]+")

	if len(phoneOnlyNumbers) != PhoneMinLength && len(phoneOnlyNumbers) != PhoneMaxLength {
		return fmt.Errorf("Invalid phone number. Use (99)99999-9999 or (99)9999-9999")
	}
	return nil
}
