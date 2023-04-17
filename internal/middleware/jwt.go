package middleware

import (
	"github.com/labstack/echo/v4"
	"magnum/internal/model/res"
	"magnum/internal/utils"
)

var whitePathList = []string{
	"/user/login",
	"/user/register",
}

// checkPath 检查path是否在白名单
func checkPath(path string) bool {
	pass := false
	for i := range whitePathList {
		if path == whitePathList[i] {
			pass = true
			break
		}
	}
	return pass
}

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Path()
		pass := checkPath(path)
		if pass {
			return next(c)
		}
		jwt, err := c.Cookie("jwt")
		if err != nil {
			return res.Unauthorized(c)
		}
		t, err := utils.ParseToken(jwt.Value)
		if err == nil {
			c.Set("user", t)
			return next(c)
		}
		return res.Unauthorized(c)
	}
}
