package v1

import (
	"github.com/labstack/echo/v4"
	"magnum/internal/model/req"
	"magnum/internal/model/res"
	"magnum/internal/service"
	"magnum/internal/utils"
	"net/http"
	"time"
)

type user struct {
	service service.User
}

var User = new(user)

type LoginResult struct {
	Valid   bool   `json:"valid"`
	Account string `json:"account"`
	Name    string `json:"name"`
	ID      string `json:"id"`
}

func (api user) Login(c echo.Context) error {
	form := new(req.Login)
	if err := utils.Bind(c, form); err != nil {
		return res.Error(c, err.Error())
	}
	token, u, err := api.service.Login(form)
	if err != nil {
		return res.Error(c, err.Error())
	}
	ck := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}
	c.SetCookie(ck)
	return res.Success(c,
		LoginResult{
			ID:      u.ID,
			Name:    u.Username,
			Account: u.Account,
			Valid:   true,
		})
}

func (api user) Logout(c echo.Context) error {
	ck := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		Path:     "/",
	}
	c.SetCookie(ck)
	return res.Success(c, "注销成功")
}

func (api user) ValidateToken(c echo.Context) error {
	u, err := utils.GetUserWithContext(c)
	if err != nil {
		return res.Error(c, "token失效")
	}
	return res.Success(c, LoginResult{
		ID:      u.ID,
		Name:    u.Username,
		Account: u.Account,
		Valid:   true,
	})
}

func (api user) Register(c echo.Context) error {
	form := new(req.Register)
	if err := utils.Bind(c, form); err != nil {
		return res.Error(c, err.Error())
	}
	if ru, err := api.service.Register(form); err != nil {
		return res.Error(c, err.Error())
	} else {
		token, err := utils.GenerateToken(ru)
		if err != nil {
			return res.Error(c, err.Error())
		}
		ck := &http.Cookie{
			Name:     "jwt",
			Value:    token,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		}
		c.SetCookie(ck)
		return res.Success(c, ru)
	}
}
