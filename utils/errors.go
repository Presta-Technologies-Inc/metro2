package utils

import (
	"errors"
	"fmt"
)

var (
	//ErrNonAlphanumeric is given when a field has non-alphanumeric characters
	ErrNonAlphanumeric = errors.New("has non alphanumeric characters")
	//ErrUpperAlpha is given when a field is not in uppercase
	ErrUpperAlpha = errors.New("is not uppercase A-Z or 0-9")
	//ErrFieldRequired is given when a field is required
	ErrFieldRequired = errors.New("is a required field")
	//ErrUpperAlpha is given when a field is not numeric characters
	ErrNumeric = errors.New("is not numeric characters")
	//ErrUpperAlpha is given when a field is an invalid phone number
	ErrPhoneNumber = errors.New("is an invalid phone number")
	//ErrValidYear is given when there's an invalid date
	ErrValidDate = errors.New("is an invalid Date")
	//ErrValidYear is given when a segment has an invalid length
	ErrSegmentLength = errors.New("has an invalid length")
	//ErrValidField is given when there's an invalid field
	ErrValidField = errors.New("is an invalid field")
)

// NewErrValidValue returns a error that has invalid value
func NewErrValidValue(field string) error {
	return fmt.Errorf("is an invalid value of %s", field)
}