package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"e.coding.net/logonod/note-server/model"
)

func (db *Database) UserCollectList(user *model.User, tag *primitive.ObjectID, start_id *primitive.ObjectID) ([]*model.UserIdTagIdsCollect, error) {
	var collects []*model.UserIdTagIdsCollect

	opts := options.Find()
	opts.SetSort(bson.D{{"_id", -1}})
	opts.SetLimit(15)

	var filter primitive.M
	if tag.Hex() == "000000000000000000000000" {
		filter = bson.M{"user_id": user.ID, "_id": bson.M{"$lt": start_id}}
	} else if tag.Hex() == "000000000000000000000001" {
		filter = bson.M{"user_id": user.ID, "tag_ids": []primitive.ObjectID{}, "_id": bson.M{"$lt": start_id}}
	} else {
		filter = bson.M{"user_id": user.ID, "tag_ids": tag, "_id": bson.M{"$lt": start_id}}
	}

	cursor, err := db.Database("note").Collection("userid_tagids_collect_map").Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {

		var elem model.UserIdTagIdsCollect
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}
		elem.Model.UpdatedAt = nil
		if *elem.UserCollected == 0 {
			elem.UserCollected = nil
		}

		collects = append(collects, &elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if collects == nil {
		collects = make([]*model.UserIdTagIdsCollect, 0)
	}

	return collects, nil
}

func (db *Database) UserCollectCreate(user *model.User, tagIds []*primitive.ObjectID, title string, desc string, url string, urlHash string, domain string) (*model.UserIdTagIdsCollect, error) {
	var collect model.UserIdTagIdsCollect
	var err error

	opts := options.FindOneAndUpdateOptions{}
	opts.SetUpsert(true)

	now := time.Now()

	if len(tagIds) == 0 {
		err = db.Database("note").Collection("userid_tagids_collect_map").FindOneAndUpdate(context.TODO(), bson.M{"user_id": user.ID, "url_hash": urlHash}, bson.M{"$setOnInsert": bson.M{"type": "webpage", "user_id": user.ID, "title": title, "cover": "", "description": desc, "url": url, "url_hash": urlHash, "site_domain": domain, "crawl_status": "pending", "user_collected": 0, "created_at": now}, "$set": bson.M{"tag_ids": []primitive.ObjectID{}, "updated_at": now}}, &opts).Decode(&collect)
		if err != nil && err != mongo.ErrNoDocuments {
			return nil, err
		}
	} else {
		err = db.Database("note").Collection("userid_tagids_collect_map").FindOneAndUpdate(context.TODO(), bson.M{"user_id": user.ID, "url_hash": urlHash}, bson.M{"$setOnInsert": bson.M{"type": "webpage", "user_id": user.ID, "title": title, "cover": "", "description": desc, "url": url, "url_hash": urlHash, "site_domain": domain, "crawl_status": "pending", "user_collected": 0, "created_at": now}, "$addToSet": bson.M{"tag_ids": bson.M{"$each": tagIds}}, "$set": bson.M{"updated_at": now}}, &opts).Decode(&collect)
		if err != nil && err != mongo.ErrNoDocuments {
			return nil, err
		}
	}

	if err == mongo.ErrNoDocuments {
		s := "webpage"
		collect.Type = &s
		collect.Title = &title
		collect.Description = &desc
		collect.Url = &url
		collect.UrlHash = &urlHash
		collect.SiteDomain = &domain
		collect.CreatedAt = &now
		collect.UpdatedAt = &now
		collect.UserCollected = nil
	}

	return &collect, nil
}

func (db *Database) UserCollectDelete(user *model.User, collectId string, tagId *primitive.ObjectID) (*bool, error) {
	tagIdHex := tagId.Hex()
	if tagIdHex == "000000000000000000000000" {
		r, err := db.Database("note").Collection("userid_tagids_collect_map").DeleteOne(context.TODO(), bson.M{"user_id": user.ID, "url_hash": collectId})
		if err != nil {
			return nil, err
		}
		if r.DeletedCount == 0 {
			m := false
			return &m, nil
		}

		m := true
		return &m, nil
	} else if tagIdHex == "000000000000000000000001" {
		r, err := db.Database("note").Collection("userid_tagids_collect_map").DeleteOne(context.TODO(), bson.M{"user_id": user.ID, "url_hash": collectId, "tag_ids": []primitive.ObjectID{}})
		if err != nil {
			return nil, err
		}
		if r.DeletedCount == 0 {
			m := false
			return &m, nil
		}

		m := true
		return &m, nil
	} else {
		rd, err := db.Database("note").Collection("userid_tagids_collect_map").DeleteOne(context.TODO(), bson.M{"user_id": user.ID, "url_hash": collectId, "tag_ids": [1]primitive.ObjectID{*tagId}})
		if err != nil {
			return nil, err
		}

		now := time.Now()
		ru, err := db.Database("note").Collection("userid_tagids_collect_map").UpdateOne(context.TODO(), bson.M{"user_id": user.ID, "url_hash": collectId, "tag_ids": tagId}, bson.M{"$pull": bson.M{"tag_ids": tagId}, "$set": bson.M{"updated_at": now}})
		if err != nil {
			return nil, err
		}

		if rd.DeletedCount == 0 && ru.MatchedCount == 0 {
			m := false
			return &m, nil
		}

		m := true
		return &m, nil
	}
}

func (db *Database) UserCollectUpdate(user *model.User, collectId string, fromTagId *primitive.ObjectID, toTagIds []*primitive.ObjectID) (*bool, error) {
	fromTagIdHex := fromTagId.Hex()
	var filter primitive.M
	if fromTagIdHex == "000000000000000000000000" {
		filter = bson.M{"user_id": user.ID, "url_hash": collectId}
	} else if fromTagIdHex == "000000000000000000000001" {
		filter = bson.M{"user_id": user.ID, "url_hash": collectId, "tag_ids": []primitive.ObjectID{}}
	} else {
		filter = bson.M{"user_id": user.ID, "url_hash": collectId, "tag_ids": fromTagId}
	}

	var writes []mongo.WriteModel
	now := time.Now()
	if len(toTagIds) == 0 {
		update := bson.M{"$set": bson.M{"tag_ids": toTagIds, "updated_at": now}}
		model := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update)
		writes = append(writes, model)
	} else {
		update1 := bson.M{"$pull": bson.M{"tag_ids": fromTagId}}
		filter2 := bson.M{"user_id": user.ID, "url_hash": collectId}
		update2 := bson.M{"$addToSet": bson.M{"tag_ids": toTagIds[0]}, "$set": bson.M{"updated_at": now}}
		model1 := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update1)
		model2 := mongo.NewUpdateOneModel().SetFilter(filter2).SetUpdate(update2)
		writes = append(writes, model1)
		writes = append(writes, model2)
	}

	r, err := db.Database("note").Collection("userid_tagids_collect_map").BulkWrite(context.TODO(), writes)
	if err != nil {
		return nil, err
	}

	if r.MatchedCount == 0 {
		m := false
		return &m, nil
	}

	m := true
	return &m, nil
}

func (db *Database) IsUserCollectExist(user *model.User, urlHash string, tagIds []*primitive.ObjectID) (bool, error) {
	var collect model.UserIdTagIdsCollect

	err := db.Database("note").Collection("userid_tagids_collect_map").FindOne(context.TODO(), bson.M{"user_id": user.ID, "url_hash": urlHash}).Decode(&collect)
	if err == mongo.ErrNoDocuments {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return model.IdsContainsIds(collect.Tags, tagIds), nil
}

func (db *Database) GetCollectIdFromTagId(user *model.User, tag *primitive.ObjectID) (*primitive.ObjectID, error) {
	var filter primitive.M
	if tag.Hex() == "000000000000000000000001" {
		filter = bson.M{"user_id": user.ID, "tag_ids": []primitive.ObjectID{}}
	} else {
		filter = bson.M{"user_id": user.ID, "tag_ids": tag}
	}

	options := options.FindOne()
	options.SetProjection(bson.D{{"_id", 1}})

	var collect model.UserIdTagIdsCollect

	err := db.Database("note").Collection("userid_tagids_collect_map").FindOne(context.TODO(), filter, options).Decode(&collect)
	if err != nil {
		return nil, err
	}

	return collect.ID, nil
}
