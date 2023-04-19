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
	all, err := api.service.ListAllByUid(uid)
	if err != nil {
		return res.Error(c, "获取表单列表失败")
	}
	return res.Success(c, all)
}

func (api form) Release(c echo.Context) error {
	rf := new(req.ReleaseForm)
	if err := utils.Bind(c, rf); err != nil {
		return res.Error(c, "表单数据格式错误，请检查。"+err.Error())
	}
	isRelease, err := api.service.QueryRelease(rf.FormID)
	if err != nil {
		return res.Error(c, "发布表单失败:"+err.Error())
	}
	if !isRelease {
		ret, err := api.service.ReleaseForm(rf)
		if err != nil {
			return res.Error(c, "发布表单失败："+err.Error())
		}
		return res.Msg(c, ret, "发布表单成功")
	}
	ret, err := api.service.UpdateForm(rf)
	if err != nil {
		return res.Error(c, "更新表单失败："+err.Error())
	}
	return res.Msg(c, ret, "更新表单成功")

}

type infoReq struct {
	ID string `json:"id"`
}

func (api form) Template(c echo.Context) error {
	f := new(infoReq)
	if err := utils.Bind(c, f); err != nil {
		return res.Error(c, "id错误")
	}
	fi, err := api.service.FindById(f.ID)
	if err != nil {
		return res.Error(c, "获取表单模板失败:"+err.Error())
	}
	t, err := api.service.FindFormTemplateById(f.ID)
	if err != nil {
		return res.Error(c, "获取表单模板失败:"+err.Error())
	}
	ret := &res.FormTemplate{
		Form:     fi,
		Template: t,
	}
	return res.Success(c, ret)
}
