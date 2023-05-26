package api

import (
	"encoding/json"
	"net/http"

	"e.coding.net/logonod/note-server/app"
	"e.coding.net/logonod/note-server/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserCollectListResponse struct {
	Collects      []*model.UserIdTagIdsCollect `json:"collects"`
	NextCollectId string                       `json:"next_collectid"`
}

func (a *API) UserCollectList(ctx *app.Context, w http.ResponseWriter, r *http.Request) *app.AppError {
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

	inputNext := r.URL.Query().Get("next_collectid")
	inputTagName := r.URL.Query().Get("tag_name")

	collects, err1 := ctx.UserCollectList(user, inputTagName, inputNext)
	if err1 != nil {
		return err1
	}

	var next string
	if len(collects) == 15 {
		next = (*collects[len(collects)-1].ID).Hex()
	}

	data, err := json.Marshal(&UserCollectListResponse{Collects: collects, NextCollectId: next})
	if err != nil {
		return &app.AppError{err, "输出错误", 40001}
	}

	if _, err = w.Write(data); err != nil {
		return &app.AppError{err, "内部错误", 500}
	}

	return nil
}

type UserCollectCreateResponse struct {
	Collect *model.UserIdTagIdsCollect `json:"collect"`
}

func (a *API) UserCollectCreate(ctx *app.Context, w http.ResponseWriter, r *http.Request) *app.AppError {
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

	inputTagNames := r.URL.Query()["n"]
	inputUrl := r.URL.Query().Get("u")
	inputTitle := r.URL.Query().Get("t")
	inputDesc := r.URL.Query().Get("d")

	collect, err1 := ctx.UserCollectCreate(user, inputTagNames, inputTitle, inputDesc, inputUrl)
	if err1 != nil {
		return err1
	}

	data, err := json.Marshal(&UserCollectCreateResponse{Collect: collect})
	if err != nil {
		return &app.AppError{err, "输出错误", 40001}
	}

	if _, err = w.Write(data); err != nil {
		return &app.AppError{err, "内部错误", 500}
	}

	return nil
}

func (a *API) UserCollectDelete(ctx *app.Context, w http.ResponseWriter, r *http.Request) *app.AppError {
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

	inputCollectId := r.URL.Query().Get("collect_id")
	inputTagName := r.URL.Query().Get("tag_name")

	err1 := ctx.UserCollectDelete(user, inputCollectId, inputTagName)
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

func (a *API) UserCollectUpdate(ctx *app.Context, w http.ResponseWriter, r *http.Request) *app.AppError {
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

	inputCollectId := r.URL.Query().Get("collect_id")
	inputFromTagName := r.URL.Query().Get("from_tag_name")
	inputToTagName := r.URL.Query().Get("to_tag_name")

	err1 := ctx.UserCollectUpdate(user, inputCollectId, inputFromTagName, inputToTagName)
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

func (a *API) UserCollectSearch(ctx *app.Context, w http.ResponseWriter, r *http.Request) *app.AppError {
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

	inputPage := r.URL.Query().Get("page")
	inputKeyword := r.URL.Query().Get("keyword")

	collects, err1 := ctx.UserCollectSearch(user, inputKeyword, inputPage)
	if err1 != nil {
		return err1
	}

	data, err := json.Marshal(&UserCollectListResponse{Collects: collects})
	if err != nil {
		return &app.AppError{err, "输出错误", 40001}
	}

	if _, err = w.Write(data); err != nil {
		return &app.AppError{err, "内部错误", 500}
	}

	return nil
}
