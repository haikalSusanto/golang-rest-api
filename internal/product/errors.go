package product

import "github.com/pkg/errors"

var (
	ErrInputValidation = errors.New("input validation error")
	ErrInternalServer  = errors.New("internal server error")
	ErrNotFound        = errors.New("records not found")
)
