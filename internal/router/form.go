package router

import (
	"github.com/labstack/echo/v4"
	v1 "magnum/api/v1"
)

type Form struct {
}

var form = new(Form)

func (f Form) Setup(router *echo.Echo) {
	r := router.Group("/form")
	r.POST("/new", v1.Form.New)
	r.POST("/list-all", v1.Form.ListAll)
	r.POST("/release", v1.Form.Release)
	r.POST("/temp", v1.Form.Template)
}
