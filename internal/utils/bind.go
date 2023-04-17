package utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func Bind[T any](c echo.Context, item T) error {
	if err := c.Bind(item); err != nil {
		return fmt.Errorf("绑定失败:%w", err)
	}
	if err := c.Validate(item); err != nil {
		return fmt.Errorf("参数错误:%w", err)
	}
	return nil
}
