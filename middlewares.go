package gofnd

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ContentTypeMiddleware(expectedContentType string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ct := ParseContentType(ctx.Request().Header)

			if ct == expectedContentType {
				return next(ctx)
			}

			return echo.NewHTTPError(http.StatusUnsupportedMediaType, "Unsupported Media Type")
		}
	}
}
