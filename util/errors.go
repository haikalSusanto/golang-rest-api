package util

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func ErrorWrapWithContext(ctx echo.Context, statusCode int, err error, additionalMessage ...string) error {
	var message []string
	for _, i := range additionalMessage {
		message = append(message, i)
	}

	ctx.Set("errorCode", statusCode)
	if len(message) != 0 {
		return errors.Wrap(err, strings.Join(message, ";"))
	}

	return err
}
