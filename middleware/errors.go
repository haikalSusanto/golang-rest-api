package middleware

import "github.com/pkg/errors"

var (
	ErrNotFound = errors.New("records not found")
)
