package app

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/mailru/easyjson"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	urlmodule "net/url"
	"strconv"
	"strings"
	"time"

	"e.coding.net/logonod/note-server/model"
)

func (ctx *Context) UserCollectList(user *model.User, inputTagName string, inputNext string) ([]*model.UserIdTagIdsCollect, *AppError) {
	nextCollectId, err := GetStartIdFromString(inputNext)
	if err != nil {
		return nil, &AppError{err, "输入错误", 40000}
	}

	var tagId *primitive.ObjectID

	inputTagName = strings.TrimSpace(inputTagName)

	if len(inputTagName) > 0 {
		tagId, err = ctx.GetUserTagIdFromTagName(user, inputTagName)

		if err == mongo.ErrNoDocuments {
			collects := make([]*model.UserIdTagIdsCollect, 0)
			return collects, nil
		}

		if err != nil {
			return nil, &AppError{err, "请稍后重试", 40002}
		}
	}

	if tagId == nil {
		i, _ := primitive.ObjectIDFromHex("000000000000000000000000")
		tagId = &i
	}

	collects, err := ctx.Database.UserCollectList(user, tagId, nextCollectId)
	if err == mongo.ErrNoDocuments {
		collects := make([]*model.UserIdTagIdsCollect, 0)
		return collects, nil
	}
	if err != nil {
		return nil, &AppError{err, "请稍后重试", 40002}
	}

	return collects, nil
}

func (ctx *Context) GetUserTagIdFromTagName(user *model.User, input string) (*primitive.ObjectID, error) {
	if input == "全部" {
		id, _ := primitive.ObjectIDFromHex("000000000000000000000000")
		return &id, nil
	}

	if input == "稍后阅读" {
		id, _ := primitive.ObjectIDFromHex("000000000000000000000001")
		return &id, nil
	}

	id, err := ctx.Database.GetUserTagIdFromTagName(user, input)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func GetStartIdFromString(input string) (*primitive.ObjectID, error) {
	if input == "" {
		input = "ffffffffffffffffffffffff"
	}

	next, err := primitive.ObjectIDFromHex(input)
	if err != nil {
		return nil, err
	}

	return &next, nil
}

func (ctx *Context) UserCollectCreate(user *model.User, inputTagNames []string, inputTitle string, inputDesc string, inputUrl string) (*model.UserIdTagIdsCollect, *AppError) {
	// 标题处理
	title := strings.TrimSpace(inputTitle)
	if len(title) == 0 {
		return nil, &AppError{errors.New(""), "标题不能为空", 40300}
	}
	titleEncode := urlmodule.QueryEscape(title)
	if len(titleEncode) > 300 {
		return nil, &AppError{errors.New(""), "标题过长", 40301}
	}

	// url处理
	url := strings.TrimSpace(inputUrl)
	if len(url) == 0 {
		return nil, &AppError{errors.New(""), "网址不能为空", 40302}
	}
	urlEncode := urlmodule.QueryEscape(url)
	if len(urlEncode) > 1200 {
		return nil, &AppError{errors.New(""), "网址过长", 40303}
	}
	u, err := urlmodule.Parse(url)
	if err != nil {
		return nil, &AppError{err, "输入错误", 40000}
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, &AppError{errors.New(""), "不支持的网页类型", 40304}
	}
	if len(u.Host) == 0 {
		return nil, &AppError{errors.New(""), "输入链接错误", 40305}
	}
	urlHash := GetUrlHash(url)

	// 描述信息的处理
	desc := strings.TrimSpace(inputDesc)
	descEncode := urlmodule.QueryEscape(desc)
	if len(descEncode) > 320 {
		return nil, &AppError{errors.New(""), "描述过长", 40306}
	}

	// 插入到延时队列中
	now := time.Now()
	message := model.UserIdUrlHash{User: user.ID, UrlHash: &urlHash, CreatedAt: &now}
	body, err := message.Encode()
	if err != nil {
		return nil, &AppError{err, "输入错误", 40000}
	}
	err = ctx.Mq.UserCollectCreate(body)
	if err != nil {
		return nil, &AppError{err, "请稍后重试", 40002}
	}

	// 分组处理
	var tagIds []*primitive.ObjectID

	if len(inputTagNames) > 0 {
		if len(inputTagNames) > 3 {
			return nil, &AppError{err, "分组最多为3个", 40307}
		}

		for i, elem := range inputTagNames {
			elemTrim := strings.TrimSpace(elem)
			inputTagNames[i] = elemTrim
			if len(elemTrim) > 15 {
				return nil, &AppError{err, "分组名称过长", 40308}
			}

			if elemTrim == "全部" {
				return nil, &AppError{err, "输入错误", 40000}
			}

			if elemTrim == "稍后阅读" && len(inputTagNames) > 1 {
				return nil, &AppError{err, "输入错误", 40000}
			}
		}

		if len(inputTagNames) == 1 && inputTagNames[0] == "稍后阅读" {
			tagIds = make([]*primitive.ObjectID, 0)
		} else {
			// 插入分组并根据tagNames获取tagIds
			tagIds, err = ctx.Database.InsertAndFindTagIdsFromTagNames(user, inputTagNames)
			if err != nil {
				return nil, &AppError{err, "请稍后重试", 40002}
			}
		}
	} else {
		tagIds = make([]*primitive.ObjectID, 0)
	}

	// 查看是否已存在
	exist, err := ctx.Database.IsUserCollectExist(user, urlHash, tagIds)
	if err != nil {
		return nil, &AppError{err, "请稍后重试", 40002}
	}
	if exist {
		return nil, &AppError{errors.New(""), "收藏已存在", 400309}
	}

	// 插入网页链接
	collect, err := ctx.Database.UserCollectCreate(user, tagIds, title, desc, url, urlHash, u.Host)
	if err != nil {
		return nil, &AppError{err, "请稍后重试", 40002}
	}

	return collect, nil
}

func (ctx *Context) UserCollectDelete(user *model.User, inputCollectId string, inputTagName string) *AppError {
	inputCollectId = strings.TrimSpace(inputCollectId)
	if len(inputCollectId) != 42 {
		return &AppError{errors.New(""), "收藏id不合法", 40310}
	}

	inputTagName = strings.TrimSpace(inputTagName)
	if len(inputTagName) > 15 {
		return &AppError{errors.New(""), "分组名称过长", 40311}
	}
	if len(inputTagName) == 0 {
		return &AppError{errors.New(""), "分组名称不能为空", 40312}
	}

	tagId, err := ctx.GetUserTagIdFromTagName(user, inputTagName)
	if err == mongo.ErrNoDocuments {
		return &AppError{err, "分组不存在", 40313}
	}
	if err != nil {
		return &AppError{err, "请稍后重试", 40002}
	}

	m, err := ctx.Database.UserCollectDelete(user, inputCollectId, tagId)
	if err != nil {
		return &AppError{err, "请稍后重试", 40002}
	}
	if !(*m) {
		return &AppError{err, "收藏不存在", 40314}
	}

	return nil
}

func (ctx *Context) UserCollectUpdate(user *model.User, inputCollectId string, inputFromTagName string, inputToTagName string) *AppError {
	inputCollectId = strings.TrimSpace(inputCollectId)
	inputFromTagName = strings.TrimSpace(inputFromTagName)
	inputToTagName = strings.TrimSpace(inputToTagName)
	if len(inputFromTagName) > 15 || len(inputToTagName) > 15 {
		return &AppError{errors.New(""), "分组名称过长", 40320}
	}
	if len(inputFromTagName) == 0 || len(inputToTagName) == 0 {
		return &AppError{errors.New(""), "分组名称不能为空", 40321}
	}
	if len(inputCollectId) != 42 {
		return &AppError{errors.New(""), "收藏id不合法", 40322}
	}
	if inputToTagName == "全部" {
		return &AppError{errors.New(""), "收藏已在全部分组中", 40323}
	}

	var toTagIds []*primitive.ObjectID
	if inputToTagName == "稍后阅读" {
		toTagIds = make([]*primitive.ObjectID, 0)
	} else {
		ids, err := ctx.Database.InsertAndFindTagIdsFromTagNames(user, []string{inputToTagName})
		if err != nil {
			return &AppError{err, "请稍后重试", 40002}
		}
		toTagIds = ids
	}

	exist, err := ctx.Database.IsUserCollectExist(user, inputCollectId, toTagIds)
	if err != nil {
		return &AppError{err, "请稍后重试", 40002}
	}
	if exist {
		return &AppError{errors.New(""), "收藏已在分组中", 400324}
	}

	var fromTagId *primitive.ObjectID
	if inputFromTagName == "稍后阅读" {
		id, _ := primitive.ObjectIDFromHex("000000000000000000000001")
		fromTagId = &id
	} else {
		id, err := ctx.GetUserTagIdFromTagName(user, inputFromTagName)
		fromTagId = id
		if err == mongo.ErrNoDocuments {
			return &AppError{err, "分组不存在", 40313}
		}
		if err != nil {
			return &AppError{err, "请稍后重试", 40002}
		}
	}

	m, err := ctx.Database.UserCollectUpdate(user, inputCollectId, fromTagId, toTagIds)
	if err != nil {
		return &AppError{err, "请稍后重试", 40002}
	}
	if !(*m) {
		return &AppError{err, "收藏不存在", 40325}
	}

	return nil
}

func (ctx *Context) UserCollectSearch(user *model.User, inputKeyword string, inputPage string) ([]*model.UserIdTagIdsCollect, *AppError) {
	page, err := strconv.Atoi(inputPage)
	if err != nil {
		fmt.Println("Error during conversion")
		return nil, &AppError{err, "输入错误", 40000}
	}

	inputKeyword = strings.TrimSpace(inputKeyword)

	if len(inputKeyword) > 0 {
		es, err := elasticsearch7.NewDefaultClient()
		if err != nil {
			return nil, &AppError{err, "请稍后重试", 40002}
		}

		str := fmt.Sprintf(`{
			"from": %d,
    		"size": 10,
			"query": { 
				"bool": {
					"should": [
						{ "match": { "title": "%s" }},
						{ "match": { "full_text": "%s" }},
						{ "match": { "meta_description": "%s" }},
						{ "match": { "url": "%s" }}
					]
				}
			},
			"highlight": {
				"number_of_fragments": 1,
				"fragment_size": 100,
				"require_field_match": "true",
				"fields": [
				  {
					"title": {
					  "number_of_fragments":1
					}
				  },
				  {
					"full_text": {
					  "number_of_fragments":1
					}
				  }
				]
			}
		}`, page*10, inputKeyword, inputKeyword, inputKeyword, inputKeyword)

		res, err := es.Search(
			es.Search.WithContext(context.Background()),
			es.Search.WithIndex("note__userid_tagids_collect_map"),
			es.Search.WithBody(strings.NewReader(str)),
			es.Search.WithTrackTotalHits(true),
			es.Search.WithPretty(),
		)
		if err != nil {
			return nil, &AppError{err, "请稍后重试", 40002}
		}
		defer res.Body.Close()
		if res.IsError() {
			var e map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
				return nil, &AppError{err, "请稍后重试", 40002}
			} else {
				return nil, &AppError{err, "请稍后重试", 40002}
			}
		}

		var ir model.SearchCollectResponse
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, &AppError{err, "请稍后重试", 40002}
		}
		err = easyjson.Unmarshal(body, &ir)
		if err != nil {
			return nil, &AppError{err, "请稍后重试", 40002}
		}
		var collects []*model.UserIdTagIdsCollect
		for _, v := range ir.Hits.Hits {
			if v.Highlight.Fulltext != nil {
				v.Source.Description = &v.Highlight.Fulltext[0]
			}
			if v.Highlight.Title != nil {
				v.Source.Title = &v.Highlight.Title[0]
			}
			collects = append(collects, &v.Source)
		}
		return collects, nil
	}

	collects := make([]*model.UserIdTagIdsCollect, 0)

	return collects, nil
}

func GetUrlHash(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	hash := hex.EncodeToString(h.Sum(nil))

	return hash + strconv.Itoa(len(input))
}
