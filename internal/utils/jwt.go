package utils

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"magnum/ent"
	"magnum/internal/model"
	"time"
)

var JWTKey = []byte("xenoblade")

type LoginClaims struct {
	Id       int64  `json:"id"`
	Account  string `json:"account"`
	Username string `json:"username"`
	Role     int64  `json:"role"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(user *ent.User) (string, error) {
	claims := entity.Claims{
		ID:       user.ID,
		Account:  user.Account,
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Unix() + 3600*24*1,
			Id:        "",
			IssuedAt:  time.Now().Unix(),
			Issuer:    user.Account,
			NotBefore: 0,
			Subject:   "login",
		},
	}
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := withClaims.SignedString(JWTKey)
	if err != nil {
		return "", err
	}
	return t, nil
}

// GetUserWithContext 从上下文获取用户信息
func GetUserWithContext(c echo.Context) (*entity.Claims, error) {
	info := c.Get("user")
	user, ok := info.(*entity.Claims)
	if !ok {
		return nil, errors.New("用户信息不合法")
	}
	return user, nil
}

// ParseToken 解析token
func ParseToken(token string) (*entity.Claims, error) {
	t, err := jwt.ParseWithClaims(token, &entity.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := t.Claims.(*entity.Claims); ok {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
