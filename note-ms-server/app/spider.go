package app

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"

	"e.coding.net/logonod/note-ms-server/model"
)

func (ctx *Context) SpiderCollectGet(inputUser string, urlHash string) (*model.UserIdTagIdsCollect, *AppError) {
	inputUser = model.TrimWholeString(inputUser)
	if len(inputUser) != 24 {
		return nil, &AppError{nil, "输入错误", 40000}
	}
	userId, err := primitive.ObjectIDFromHex(inputUser)
	if err != nil {
		return nil, &AppError{nil, "输入错误", 40000}
	}

	urlHash = model.TrimWholeString(urlHash)
	//if len(urlHash) != 42 {
	//	return nil, &AppError{nil, "输入错误", 40000}
	//}

	collect, err := ctx.Database.SpiderCollectGet(&userId, urlHash)
	if err != nil {
		return nil, &AppError{err, "请稍后重试", 40002}
	}
	if collect == nil {
		return nil, &AppError{err, "收藏不存在", 40100}
	}

	return collect, nil
}

func (ctx *Context) SpiderWebpageGet(urlHash string) (*model.Webpage, *AppError) {
	urlHash = model.TrimWholeString(urlHash)
	if len(urlHash) != 46 {
		return nil, &AppError{nil, "输入错误", 40000}
	}

	webpage, err := ctx.Database.SpiderWebpageGet(urlHash)
	if err != nil {
		return nil, &AppError{err, "请稍后重试", 40002}
	}
	if webpage == nil {
		return nil, &AppError{err, "网页不存在", 40110}
	}

	return webpage, nil
}

func (ctx *Context) SpiderCollectCreate(inputUser string, urlHash string, title string, cover string, description string, metaDescription string, fullText string, url string, siteDomain string) *AppError {
	inputUser = model.TrimWholeString(inputUser)
	if len(inputUser) != 24 {
		return &AppError{nil, "输入错误", 40000}
	}
	userId, err := primitive.ObjectIDFromHex(inputUser)
	if err != nil {
		return &AppError{nil, "输入错误", 40000}
	}

	urlHash = model.TrimWholeString(urlHash)
	//if len(urlHash) != 42 {
	//	return &AppError{nil, "输入错误", 40000}
	//}

	title = model.TrimMutipleSpaceWholeString(title)
	if len(title) == 0 {
		return &AppError{nil, "标题不能为空", 40120}
	}
	title = model.NormSubString(title, ctx.Config.NormSubStringTitleLength)

	cover = strings.TrimSpace(cover)
	if len(cover) > 200 {
		return &AppError{nil, "缩略图链接过长", 40121}
	}

	description = model.TrimMutipleSpaceWholeString(description)
	description = model.NormSubString(description, ctx.Config.NormSubStringDescriptionLength)

	metaDescription = model.TrimMutipleSpaceWholeString(metaDescription)
	metaDescription = model.NormSubString(metaDescription, ctx.Config.NormSubStringDescriptionLength)

	fullText = model.TrimFullText(fullText)
	fullText = model.NormSubFullText(fullText, ctx.Config.NormSubFullTextLength)

	url = strings.TrimSpace(url)
	if len(url) > 4000 {
		return &AppError{nil, "链接过长", 40122}
	}

	siteDomain = model.TrimWholeString(siteDomain)
	siteDomain = model.NormSubDomain(siteDomain, ctx.Config.NormSubStringDomainLength)

	userCollected, err := ctx.Database.SpiderUserCollectedCount(urlHash)
	if err != nil {
		return &AppError{err, "请稍后重试", 40002}
	}
	if userCollected == 0 {
		return &AppError{err, "用户未收藏", 40123}
	}

	err = ctx.Database.SpiderWebpageInsertOrSetUserCollected(title, cover, description, metaDescription, fullText, url, urlHash, siteDomain, userCollected)
	if err != nil {
		return &AppError{err, "请稍后重试", 40002}
	}

	exist, err := ctx.Database.SpiderCollectCreate(&userId, urlHash, cover, description, metaDescription, fullText, userCollected)
	if err != nil {
		return &AppError{err, "请稍后重试", 40002}
	}
	if !exist {
		return &AppError{err, "用户未收藏", 40123}
	}

	return nil
}

func (ctx *Context) SpiderCollectUpdateStatus(inputUser string, urlHash string, url string, crawlStatus string) *AppError {
	inputUser = model.TrimWholeString(inputUser)
	if len(inputUser) != 24 {
		return &AppError{nil, "输入错误", 40000}
	}
	userId, err := primitive.ObjectIDFromHex(inputUser)
	if err != nil {
		return &AppError{nil, "输入错误", 40000}
	}

	urlHash = model.TrimWholeString(urlHash)
	//if len(urlHash) != 42 {
	//	return &AppError{nil, "输入错误", 40000}
	//}

	url = strings.TrimSpace(url)
	if len(url) > 4000 {
		return &AppError{nil, "输入错误", 40000}
	}
	description := model.TrimMutipleSpaceWholeString(url)
	description = model.NormSubString(description, ctx.Config.NormSubStringDescriptionLength)

	crawlStatus = model.TrimWholeString(inputUser)
	if crawlStatus != "pending" && crawlStatus != "finished" && crawlStatus != "failed" && crawlStatus != "unsupported" {
		return &AppError{nil, "输入错误", 40000}
	}

	exist, err := ctx.Database.SpiderCollectUpdateStatus(&userId, urlHash, description, crawlStatus)
	if err != nil {
		return &AppError{err, "请稍后重试", 40002}
	}
	if !exist {
		return &AppError{err, "用户未收藏", 40130}
	}

	return nil
}
