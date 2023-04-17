package service

import (
	"context"
	"errors"
	"fmt"
	"magnum/ent"
	"magnum/ent/user"
	"magnum/internal/app"
	"magnum/internal/model/req"
	"magnum/internal/utils"
)

type User struct {
}

func (User) Login(f *req.Login) (string, *ent.User, error) {
	u, err := app.DB.User.Query().
		Where(user.Username(f.Username), user.Password(f.Password)).
		Only(context.Background())
	if err != nil {
		return "", nil, errors.New("登录失败，请检查用户名和密码")
	}
	token, err := utils.GenerateToken(u)
	if err != nil {
		return "", nil, fmt.Errorf("生成token失败：%w", err)
	}
	return token, u, nil
}

func (User) Register(f *req.Register) (*ent.User, error) {
	only, _ := app.DB.User.Query().Where(user.Username(f.Username)).Only(context.Background())
	if only != nil {
		return nil, errors.New("用户已注册")
	}
	return app.DB.User.Create().
		SetUsername(f.Username).
		SetPassword(f.Password).
		SetAccount(f.Username).
		SetRole("user").
		SetPhone("").
		Save(context.Background())
}
