package gofnd

import (
	"fmt"
	"mime"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ParseContentType(header http.Header) string {
	contentTypeRaw, ok := header["Content-Type"]
	if !ok {
		return ""
	}
	contentType, _, _ := mime.ParseMediaType(contentTypeRaw[0])

	return contentType
}

func BindAndValidate(ctx echo.Context, body interface{}) (err error) {
	if err = ctx.Bind(body); err != nil {
		return fmt.Errorf("gofnd[.BindAndValidate][1]: %+v", err)
	}

	if err = ctx.Validate(body); err != nil {
		return fmt.Errorf("gofnd[.BindAndValidate][2]: %+v", err)
	}

	return
}
