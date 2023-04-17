package app

import (
	"github.com/labstack/echo/v4"
	"magnum/ent"
)

var Router *echo.Echo
var DB *ent.Client
