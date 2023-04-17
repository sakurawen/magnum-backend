package v1

import (
	"github.com/labstack/echo/v4"
	"magnum/internal/model/req"
	"magnum/internal/model/res"
	"magnum/internal/service"
	"magnum/internal/utils"
)

type form struct {
	service service.Form
}

var Form = new(form)

func (api form) New(c echo.Context) error {
	f := new(req.NewForm)
	if err := utils.Bind(c, f); err != nil {
		return res.Error(c, "创建表单信息错误"+err.Error())
	}
	uc, err := utils.GetUserWithContext(c)
	ret, err := api.service.New(f, uc.ID)
	if err != nil {
		return res.Error(c, "创建表单失败"+err.Error())
	}
	return res.Success(c, ret)
}

func (api form) ListAll(c echo.Context) error {
	u, err := utils.GetUserWithContext(c)
	if err != nil {
		return res.Error(c, err.Error())
	}
	uid := u.ID
	all, err := api.service.ListAll(uid)
	if err != nil {
		return res.Error(c, "获取表单列表失败")
	}
	return res.Success(c, all)
}
