package gofnd

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewNotFoundHTTPError(msg string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusNotFound, msg)
}

// NewBadRequestHTTPError : if arg == "" -> use default message "bad request"
func NewBadRequestHTTPError(msg string, err error) *echo.HTTPError {
	m := "bad request"

	if msg != "" {
		m += fmt.Sprintf(": %v", msg)
	}

	return echo.NewHTTPError(http.StatusBadRequest, m).SetInternal(err)
}

func NewInternalServerError(err error) *echo.HTTPError {
	return echo.ErrInternalServerError.SetInternal(err)
}

func AlreadyExistError(msg string, err error) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusBadRequest, msg).SetInternal(err)
}
