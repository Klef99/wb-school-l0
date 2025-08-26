package models

import (
	"fmt"
)

type ValidationError struct {
	Field   string
	Value   any
	Message string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("Validation failed on field '%s' with value '%v': %s", v.Field, v.Value, v.Message)
}
