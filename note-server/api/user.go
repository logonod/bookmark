package api

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"time"

	"e.coding.net/logonod/note-server/app"
	"e.coding.net/logonod/note-server/model"
)

type UserLoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Status string `json:"status"`
}

func (a *API) UserLogin(ctx *app.Context, w http.ResponseWriter, r *http.Request) *app.AppError {
	var input UserLoginRequest

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &app.AppError{err, "输入错误", 40000}
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return &app.AppError{err, "输入错误", 40000}
	}

	user := &model.User{Phone: input.Phone}
	if err := ctx.UserLogin(user, input.Password); err != nil {
		return err
	}

	data, err := json.Marshal(&UserLoginResponse{Status: "ok"})
	if err != nil {
		return &app.AppError{err, "输出错误", 40001}
	}

	uid := uuid.New().String()

	if err := ctx.SetUserCookie(uid, user, a.Config.SessionTTL); err != nil {
		return &app.AppError{err, "用户登陆失败", 40103}
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "uid",
		Value:   uid,
		Expires: time.Now().Add(time.Duration(a.Config.SessionTTL) * time.Second),
		Path:    "/",
	})

	if _, err = w.Write(data); err != nil {
		return &app.AppError{err, "内部错误", 500}
	}

	return nil
}

func (a *API) UserLogout(ctx *app.Context, w http.ResponseWriter, r *http.Request) *app.AppError {
	cookie, err := r.Cookie("uid")
	if err != nil {
		if err == http.ErrNoCookie {
			return &app.AppError{err, "未登陆", 40003}
		}
		return &app.AppError{err, "输入错误", 40000}
	}

	_, err = ctx.GetUserByCookie(cookie)
	if err != nil {
		if err == redis.Nil {
			return &app.AppError{err, "未登陆", 40003}
		}
		return &app.AppError{err, "请稍后重试", 40002}
	}

	data, err := json.Marshal(&UserLoginResponse{Status: "ok"})
	if err != nil {
		return &app.AppError{err, "输出错误", 40001}
	}

	if err := ctx.DeleteUserCookie(cookie.Value); err != nil {
		return &app.AppError{err, "用户登出失败", 40110}
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "uid",
		Value:   "",
		Expires: time.Now().Add(time.Duration(-a.Config.SessionTTL) * time.Second),
		Path:    "/",
	})

	if _, err = w.Write(data); err != nil {
		return &app.AppError{err, "内部错误", 500}
	}

	return nil
}
