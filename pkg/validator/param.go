package validator

import (
	"errors"
	"strings"
)

func ValidateParam(param, msg string) error {
	value := strings.ReplaceAll(param, " ", "")
	if len(value) == 0 {
		return errors.New(msg)
	}
	return nil
}
