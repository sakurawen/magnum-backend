package router

import (
	"github.com/labstack/echo/v4"
	v1 "magnum/api/v1"
)

type Form struct {
}

var form = new(Form)

func (f Form) Setup(router *echo.Echo) {
	formRouter := router.Group("/form")
	formRouter.POST("/new", v1.Form.New)
	formRouter.POST("/list-all", v1.Form.ListAll)
}
