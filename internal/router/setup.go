package router

import "github.com/labstack/echo/v4"

func Setup(router *echo.Echo) {
	user.Setup(router)
	form.Setup(router)
}
