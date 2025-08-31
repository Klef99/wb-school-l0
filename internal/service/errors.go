package service

import (
	"errors"
)

var ErrOrderNotFound = errors.New("order does not exist")

var ErrValidationFailed = errors.New("validation failed")
