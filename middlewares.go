package gofnd

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
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

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return logrusMiddlewareHandler(ctx, next)
	}
}

func CustomErrorHandler(err error, ctx echo.Context) {
	httpError := &echo.HTTPError{
		Code:    http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
	}

	if he, ok := err.(*echo.HTTPError); ok {
		httpError = he
	}

	log.Info(httpError)

	if err = ctx.JSON(httpError.Code, httpError); err != nil {
		err = fmt.Errorf("gofnd: [.CustomErrorHandler][1]: %+v", err)
		log.Error(err)
	}
}

// --- private

func logrusMiddlewareHandler(c echo.Context, next echo.HandlerFunc) error {
	req := c.Request()
	res := c.Response()
	start := time.Now()
	if err := next(c); err != nil {
		c.Error(err)
	}
	stop := time.Now()

	p := req.URL.Path

	bytesIn := req.Header.Get(echo.HeaderContentLength)

	log.WithFields(map[string]interface{}{
		"time_rfc3339":  time.Now().Format(time.RFC3339),
		"remote_ip":     c.RealIP(),
		"host":          req.Host,
		"uri":           req.RequestURI,
		"method":        req.Method,
		"path":          p,
		"referer":       req.Referer(),
		"user_agent":    req.UserAgent(),
		"status":        res.Status,
		"latency":       strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
		"latency_human": stop.Sub(start).String(),
		"bytes_in":      bytesIn,
		"bytes_out":     strconv.FormatInt(res.Size, 10),
	}).Info("Handled request")

	return nil
}
