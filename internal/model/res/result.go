package res

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type result struct {
	Code int64  `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// Success 成功响应
func Success(ctx echo.Context, data any) error {
	return ctx.JSON(http.StatusOK, &result{
		http.StatusOK,
		data,
		"",
	})
}

// Error 错误响应
func Error(ctx echo.Context, msg string) error {
	return ctx.JSON(http.StatusOK, &result{
		http.StatusBadRequest,
		nil,
		msg,
	})
}

func Unauthorized(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &result{
		http.StatusUnauthorized,
		nil,
		"权限不足",
	})

}
