package app

import (
	"context"
	"encoding/json"
	"fmt"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/mailru/easyjson"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"strings"

	"e.coding.net/logonod/note-server/model"
)

func (ctx *Context) UserTagList(user *model.User, input string) ([]*model.UserIdTag, *AppError) {
	nextTagId, err := GetStartIdFromString(input)
	if err != nil {
		return nil, &AppError{err, "输入错误", 40000}
	}

	tags, err := ctx.Database.UserTagList(user, nextTagId)
	if err != nil {
		return nil, &AppError{err, "请稍后重试", 40002}
	}

	return tags, nil
}

func (ctx *Context) UserTagCreate(user *model.User, input string) (*model.UserIdTag, *AppError) {
	name := strings.TrimSpace(input)
	if len(name) == 0 {
		return nil, &AppError{errors.New(""), "分组名称不能为空", 40200}
	}
	if len(name) > 15 {
		return nil, &AppError{errors.New(""), "分组名称过长", 40202}
	}
	if name == "全部" || name == "稍后阅读" {
		return nil, &AppError{errors.New(""), "分组已存在", 40203}
	}

	tag, err := ctx.Database.UserTagCreate(user, name)
	if err != nil {
		if writeException, ok := err.(mongo.WriteException); ok {
			if len(writeException.WriteErrors) > 0 && writeException.WriteErrors[0].Code == 11000 {
				return nil, &AppError{err, "分组已存在", 40203}
			}
		}

		return nil, &AppError{err, "请稍后重试", 40002}
	}

	return tag, nil
}

func (ctx *Context) UserTagDelete(user *model.User, inputTagName string) *AppError {
	inputTagName = strings.TrimSpace(inputTagName)
	if len(inputTagName) > 15 {
		return &AppError{errors.New(""), "分组名称过长", 40210}
	}
	if inputTagName == "全部" {
		return &AppError{errors.New(""), "输入错误", 40000}
	}
	if inputTagName == "稍后阅读" {
		return &AppError{errors.New(""), "不能删除保留分组", 40212}
	}

	var tagId *primitive.ObjectID
	if len(inputTagName) > 0 {
		id, err := ctx.Database.GetUserTagIdFromTagName(user, inputTagName)
		tagId = id
		if err == mongo.ErrNoDocuments {
			return &AppError{err, "分组不存在", 40213}
		}

		if err != nil {
			return &AppError{err, "请稍后重试", 40002}
		}
	} else {
		return &AppError{errors.New(""), "输入错误", 40000}
	}

	_, err := ctx.Database.GetCollectIdFromTagId(user, tagId)
	if err != nil && err != mongo.ErrNoDocuments {
		return &AppError{err, "请稍后重试", 40002}
	}
	if err == nil {
		return &AppError{errors.New(""), "不能删除非空分组", 40214}
	}

	r, err := ctx.Database.UserTagDelete(tagId)
	if err != nil {
		return &AppError{err, "请稍后重试", 40002}
	}
	if r.DeletedCount == 0 {
		return &AppError{err, "分组不存在", 40213}
	}

	return nil
}

func (ctx *Context) UserTagUpdate(user *model.User, inputFrom string, inputTo string) *AppError {
	inputFrom = strings.TrimSpace(inputFrom)
	inputTo = strings.TrimSpace(inputTo)
	if len(inputFrom) > 15 || len(inputTo) > 15 {
		return &AppError{errors.New(""), "分组名称过长", 40220}
	}
	if len(inputFrom) == 0 || len(inputTo) == 0 {
		return &AppError{errors.New(""), "分组名称不能为空", 40221}
	}
	if inputFrom == "全部" || inputFrom == "稍后阅读" || inputTo == "全部" || inputFrom == "稍后阅读" {
		return &AppError{errors.New(""), "不能修改保留分组", 40222}
	}

	r, err := ctx.Database.UserTagUpdate(user, inputFrom, inputTo)
	if err != nil {
		if writeException, ok := err.(mongo.WriteException); ok {
			if len(writeException.WriteErrors) > 0 && writeException.WriteErrors[0].Code == 11000 {
				return &AppError{err, "分组名称已存在", 40224}
			}
		}

		return &AppError{err, "请稍后重试", 40002}
	}
	if r.MatchedCount == 0 {
		return &AppError{err, "源分组不存在", 40223}
	}

	return nil
}

func (ctx *Context) UserTagSearch(user *model.User, inputKeyword string) ([]*model.UserIdTagSearch, *AppError) {
	//nextTagId, err := GetStartIdFromString(input)
	//if err != nil {
	//	return nil, &AppError{err, "输入错误", 40000}
	//}
	//
	//tags, err := ctx.Database.UserTagList(user, nextTagId)
	//if err != nil {
	//	return nil, &AppError{err, "请稍后重试", 40002}
	//}
	//
	//return tags, nil

	inputKeyword = strings.TrimSpace(inputKeyword)

	if len(inputKeyword) > 0 {
		es, err := elasticsearch7.NewDefaultClient()
		if err != nil {
			return nil, &AppError{err, "请稍后重试", 40002}
		}

		str := fmt.Sprintf(`{
			"from": 0,
    		"size": 100,
			"query": { 
				"bool": {
					"must": [
						{ "match": { "tag_name": "%s" }}
					]
				}
			}
		}`, inputKeyword)

		res, err := es.Search(
			es.Search.WithContext(context.Background()),
			es.Search.WithIndex("note__userid_tag_map"),
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

		var ir model.SearchTagResponse
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, &AppError{err, "请稍后重试", 40002}
		}
		err = easyjson.Unmarshal(body, &ir)
		if err != nil {
			return nil, &AppError{err, "请稍后重试", 40002}
		}
		var tags []*model.UserIdTagSearch
		for _, v := range ir.Hits.Hits {
			tags = append(tags, &v.Source)
		}
		return tags, nil
	}

	tags := make([]*model.UserIdTagSearch, 0)

	return tags, nil
}
