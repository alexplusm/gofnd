package gofnd

import (
	"fmt"

	vld "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type validator struct {
	Validator *vld.Validate
}

func NewValidator() echo.Validator {
	return &validator{Validator: vld.New()}
}

func (v *validator) Validate(i interface{}) (err error) {
	if err = v.Validator.Struct(i); err != nil {
		err = fmt.Errorf("gofnd: [validator.Validate][1]: %+v", err)
	}

	return
}
