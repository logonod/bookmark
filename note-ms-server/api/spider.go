package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"e.coding.net/logonod/note-ms-server/app"
	"e.coding.net/logonod/note-ms-server/model"
)

type SpiderCollectGetRequest struct {
	User    string `json:"user_id"`
	UrlHash string `json:"url_hash"`
}

type SpiderCollectGetResponse struct {
	Collect *model.UserIdTagIdsCollect `json:"collect"`
}

func (a *API) SpiderCollectGet(ctx *app.Context, w http.ResponseWriter, r *http.Request) *app.AppError {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &app.AppError{err, "输入错误", 40000}
	}

	var input SpiderCollectGetRequest
	if err := json.Unmarshal(body, &input); err != nil {
		return &app.AppError{err, "输入错误", 40000}
	}

	collect, err1 := ctx.SpiderCollectGet(input.User, input.UrlHash)
	if err1 != nil {
		return err1
	}

	data, err := json.Marshal(&SpiderCollectGetResponse{Collect: collect})
	if err != nil {
		return &app.AppError{err, "输出错误", 40001}
	}

	if _, err = w.Write(data); err != nil {
		return &app.AppError{err, "内部错误", 500}
	}

	return nil
}

type SpiderWebpageGetRequest struct {
	UrlHash string `json:"url_hash"`
}

type SpiderWebpageGetResponse struct {
	Webpage *model.Webpage `json:"webpage"`
}

func (a *API) SpiderWebpageGet(ctx *app.Context, w http.ResponseWriter, r *http.Request) *app.AppError {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &app.AppError{err, "输入错误", 40000}
	}

	var input SpiderWebpageGetRequest
	if err := json.Unmarshal(body, &input); err != nil {
		return &app.AppError{err, "输入错误", 40000}
	}

	webpage, err1 := ctx.SpiderWebpageGet(input.UrlHash)
	if err1 != nil {
		return err1
	}

	data, err := json.Marshal(&SpiderWebpageGetResponse{Webpage: webpage})
	if err != nil {
		return &app.AppError{err, "输出错误", 40001}
	}

	if _, err = w.Write(data); err != nil {
		return &app.AppError{err, "内部错误", 500}
	}

	return nil
}

type SpiderCollectCreateRequest struct {
	User            string `json:"user_id"`
	UrlHash         string `json:"url_hash"`
	Title           string `json:"title"`
	Cover           string `json:"cover"`
	Description     string `json:"description"`
	MetaDescription string `json:"meta_description"`
	FullText        string `json:"full_text"`
	Url             string `json:"url"`
	SiteDomain      string `json:"site_domain"`
}

type SpiderCollectCreateResponse struct {
	Status string `json:"status"`
}

func (a *API) SpiderCollectCreate(ctx *app.Context, w http.ResponseWriter, r *http.Request) *app.AppError {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &app.AppError{err, "输入错误", 40000}
	}

	var input SpiderCollectCreateRequest

	if err := json.Unmarshal(body, &input); err != nil {
		//return &app.AppError{err, "输入错误", 40000}
		return &app.AppError{err, string(body), 40000}
	}

	err1 := ctx.SpiderCollectCreate(input.User, input.UrlHash, input.Title, input.Cover, input.Description, input.MetaDescription, input.FullText, input.Url, input.SiteDomain)
	if err1 != nil {
		return err1
	}

	data, err := json.Marshal(&SpiderCollectCreateResponse{Status: "ok"})
	if err != nil {
		return &app.AppError{err, "输出错误", 40001}
	}

	if _, err = w.Write(data); err != nil {
		return &app.AppError{err, "内部错误", 500}
	}

	return nil
}

type SpiderCollectUpdateStatusRequest struct {
	User        string `json:"user_id"`
	UrlHash     string `json:"url_hash"`
	Url         string `json:"url"`
	CrawlStatus string `json:"crawl_status"`
}

type SpiderCollectUpdateStatusResponse struct {
	Status string `json:"status"`
}

func (a *API) SpiderCollectUpdateStatus(ctx *app.Context, w http.ResponseWriter, r *http.Request) *app.AppError {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &app.AppError{err, "输入错误", 40000}
	}

	var input SpiderCollectUpdateStatusRequest
	if err := json.Unmarshal(body, &input); err != nil {
		return &app.AppError{err, "输入错误", 40000}
	}

	err1 := ctx.SpiderCollectUpdateStatus(input.User, input.UrlHash, input.Url, input.CrawlStatus)
	if err1 != nil {
		return err1
	}

	data, err := json.Marshal(&SpiderCollectUpdateStatusResponse{Status: "ok"})
	if err != nil {
		return &app.AppError{err, "输出错误", 40001}
	}

	if _, err = w.Write(data); err != nil {
		return &app.AppError{err, "内部错误", 500}
	}

	return nil
}
