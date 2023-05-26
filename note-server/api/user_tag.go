package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"e.coding.net/logonod/note-server/app"
	"e.coding.net/logonod/note-server/model"
)

type UserTagListResponse struct {
	Tags      []*model.UserIdTag `json:"tags"`
	NextTagId string             `json:"next_tagid"`
}

type UserTagListSearchResponse struct {
	Tags      []*model.UserIdTagSearch `json:"tags"`
	NextTagId string                   `json:"next_tagid"`
}

type UserTagCreateResponse struct {
	Tag *model.UserIdTag `json:"tag"`
}

func (a *API) UserTagList(ctx *app.Context, w http.ResponseWriter, r *http.Request) *app.AppError {
	//cookie, err := r.Cookie("uid")
	//if err != nil {
	//	if err == http.ErrNoCookie {
	//		return &app.AppError{err, "未登陆", 40003}
	//	}
	//	return &app.AppError{err, "输入错误", 40000}
	//}

	//user, err := ctx.GetUserByCookie(cookie)
	//if err != nil {
	//	if err == redis.Nil {
	//		return &app.AppError{err, "未登陆", 40003}
	//	}
	//	return &app.AppError{err, "请稍后重试", 40002}
	//}
	id, err := primitive.ObjectIDFromHex("5d9b29abb6fe0f893f79eddc")
	user := &model.User{Model: model.Model{ID: &id}}

	input := r.URL.Query().Get("next_tagid")

	tags, err1 := ctx.UserTagList(user, input)
	if err1 != nil {
		return err1
	}

	var next string
	if len(tags) == 15 {
		next = (*tags[len(tags)-1].ID).Hex()
	}

	data, err := json.Marshal(&UserTagListResponse{Tags: tags, NextTagId: next})
	if err != nil {
		return &app.AppError{err, "输出错误", 40001}
	}

	if _, err = w.Write(data); err != nil {
		return &app.AppError{err, "内部错误", 500}
	}

	return nil
}

func (a *API) UserTagCreate(ctx *app.Context, w http.ResponseWriter, r *http.Request) *app.AppError {
	//cookie, err := r.Cookie("uid")
	//if err != nil {
	//	if err == http.ErrNoCookie {
	//		return &app.AppError{err, "未登陆", 40003}
	//	}
	//	return &app.AppError{err, "输入错误", 40000}
	//}
	//
	//user, err := ctx.GetUserByCookie(cookie)
	//if err != nil {
	//	if err == redis.Nil {
	//		return &app.AppError{err, "未登陆", 40003}
	//	}
	//	return &app.AppError{err, "请稍后重试", 40002}
	//}
	id, err := primitive.ObjectIDFromHex("5d9b29abb6fe0f893f79eddc")
	user := &model.User{Model: model.Model{ID: &id}}

	input := r.URL.Query().Get("name")

	tag, err1 := ctx.UserTagCreate(user, input)
	if err1 != nil {
		return err1
	}

	data, err := json.Marshal(&UserTagCreateResponse{Tag: tag})
	if err != nil {
		return &app.AppError{err, "输出错误", 40001}
	}

	if _, err = w.Write(data); err != nil {
		return &app.AppError{err, "内部错误", 500}
	}

	return nil
}

func (a *API) UserTagDelete(ctx *app.Context, w http.ResponseWriter, r *http.Request) *app.AppError {
	//cookie, err := r.Cookie("uid")
	//if err != nil {
	//	if err == http.ErrNoCookie {
	//		return &app.AppError{err, "未登陆", 40003}
	//	}
	//	return &app.AppError{err, "输入错误", 40000}
	//}
	//
	//user, err := ctx.GetUserByCookie(cookie)
	//if err != nil {
	//	if err == redis.Nil {
	//		return &app.AppError{err, "未登陆", 40003}
	//	}
	//	return &app.AppError{err, "请稍后重试", 40002}
	//}
	id, err := primitive.ObjectIDFromHex("5d9b29abb6fe0f893f79eddc")
	user := &model.User{Model: model.Model{ID: &id}}

	inputName := r.URL.Query().Get("name")

	err1 := ctx.UserTagDelete(user, inputName)
	if err1 != nil {
		return err1
	}

	data, err := json.Marshal(&UserLoginResponse{Status: "ok"})
	if err != nil {
		return &app.AppError{err, "输出错误", 40001}
	}

	if _, err = w.Write(data); err != nil {
		return &app.AppError{err, "内部错误", 500}
	}

	return nil
}

func (a *API) UserTagUpdate(ctx *app.Context, w http.ResponseWriter, r *http.Request) *app.AppError {
	//cookie, err := r.Cookie("uid")
	//if err != nil {
	//	if err == http.ErrNoCookie {
	//		return &app.AppError{err, "未登陆", 40003}
	//	}
	//	return &app.AppError{err, "输入错误", 40000}
	//}
	//
	//user, err := ctx.GetUserByCookie(cookie)
	//if err != nil {
	//	if err == redis.Nil {
	//		return &app.AppError{err, "未登陆", 40003}
	//	}
	//	return &app.AppError{err, "请稍后重试", 40002}
	//}
	id, err := primitive.ObjectIDFromHex("5d9b29abb6fe0f893f79eddc")
	user := &model.User{Model: model.Model{ID: &id}}

	inputFrom := r.URL.Query().Get("from")
	inputTo := r.URL.Query().Get("to")

	err1 := ctx.UserTagUpdate(user, inputFrom, inputTo)
	if err1 != nil {
		return err1
	}

	data, err := json.Marshal(&UserLoginResponse{Status: "ok"})
	if err != nil {
		return &app.AppError{err, "输出错误", 40001}
	}

	if _, err = w.Write(data); err != nil {
		return &app.AppError{err, "内部错误", 500}
	}

	return nil
}

func (a *API) UserTagSearch(ctx *app.Context, w http.ResponseWriter, r *http.Request) *app.AppError {
	//cookie, err := r.Cookie("uid")
	//if err != nil {
	//	if err == http.ErrNoCookie {
	//		return &app.AppError{err, "未登陆", 40003}
	//	}
	//	return &app.AppError{err, "输入错误", 40000}
	//}

	//user, err := ctx.GetUserByCookie(cookie)
	//if err != nil {
	//	if err == redis.Nil {
	//		return &app.AppError{err, "未登陆", 40003}
	//	}
	//	return &app.AppError{err, "请稍后重试", 40002}
	//}
	id, err := primitive.ObjectIDFromHex("5d9b29abb6fe0f893f79eddc")
	user := &model.User{Model: model.Model{ID: &id}}

	input := r.URL.Query().Get("keyword")

	tags, err1 := ctx.UserTagSearch(user, input)
	if err1 != nil {
		return err1
	}

	var next string

	data, err := json.Marshal(&UserTagListSearchResponse{Tags: tags, NextTagId: next})
	if err != nil {
		return &app.AppError{err, "输出错误", 40001}
	}

	if _, err = w.Write(data); err != nil {
		return &app.AppError{err, "内部错误", 500}
	}

	return nil
}
