package app

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"regexp"

	"e.coding.net/logonod/note-server/model"
)

func (ctx *Context) SetUserCookie(uid string, user *model.User, ttl int) error {
	return ctx.Cache.SetUserCookie(uid, user, ttl)
}

func (ctx *Context) GetUserByCookie(cookie *http.Cookie) (*model.User, error) {
	return ctx.Cache.GetUserCookie(cookie.Value)
}

func (ctx *Context) UserLogin(user *model.User, password string) *AppError {
	if err := ctx.validateUserLogin(user, password); err != nil {
		return err
	}

	err := ctx.Database.UserLogin(user)

	if err == mongo.ErrNoDocuments {
		return &AppError{err, "用户名或密码错误", 40100}
	}

	if err != nil {
		return &AppError{err, "请稍后重试", 40002}
	}

	if !user.CheckPassword(password) {
		return &AppError{err, "不正确的用户名或密码", 40100}
	}

	return nil
}

func (ctx *Context) validateUserLogin(user *model.User, password string) *AppError {
	// 简单的用户输入格式验证
	if match, _ := regexp.MatchString(`^1[3456789]\d{9}$`, user.Phone); !match {
		return &AppError{nil, "非法手机号", 40101}
	}

	if password == "" {
		return &AppError{nil, "需要输入密码", 40102}
	}

	return nil
}

func (ctx *Context) DeleteUserCookie(uid string) error {
	return ctx.Cache.DeleteUserCookie(uid)
}
