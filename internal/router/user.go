package router

import (
	"github.com/labstack/echo/v4"
	"magnum/api/v1"
)

type User struct {
}

var user = new(User)

func (u User) Setup(router *echo.Echo) {
	userRouter := router.Group("/user")
	userRouter.POST("/login", v1.User.Login)
	userRouter.POST("/logout", v1.User.Logout)
	userRouter.POST("/register", v1.User.Register)
	userRouter.POST("/validate_token", v1.User.ValidateToken)
}
