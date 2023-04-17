package initialize

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	customMiddleware "magnum/internal/middleware"
	"magnum/internal/utils"
)

func Router() *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} method=${method} uri=${uri} status=${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))
	e.Use(customMiddleware.JWT)
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	return e
}
