package errors

import "errors"

var (
	ErrAWBNotFound    = errors.New("awb not found")
	ErrAWBHasBeenUsed = errors.New("awb has been used")
	ErrAWBInvalid     = errors.New("awb is invalid")
)
