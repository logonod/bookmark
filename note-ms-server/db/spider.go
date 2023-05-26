package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"e.coding.net/logonod/note-ms-server/model"
)

func (db *Database) SpiderCollectGet(userId *primitive.ObjectID, urlHash string) (*model.UserIdTagIdsCollect, error) {
	var collect model.UserIdTagIdsCollect

	options := options.FindOne()
	options.SetProjection(bson.D{{"meta_description", 0}, {"full_text", 0}})

	err := db.Database("note").Collection("userid_tagids_collect_map").FindOne(context.TODO(), bson.M{"user_id": userId, "url_hash": urlHash}).Decode(&collect)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &collect, nil
}

func (db *Database) SpiderWebpageGet(urlHash string) (*model.Webpage, error) {
	var webpage model.Webpage

	err := db.Database("note").Collection("webpage").FindOne(context.TODO(), bson.M{"url_hash": urlHash}).Decode(&webpage)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &webpage, nil
}

func (db *Database) SpiderUserCollectedCount(urlHash string) (int64, error) {
	count, err := db.Database("note").Collection("userid_tagids_collect_map").CountDocuments(context.TODO(), bson.M{"url_hash": urlHash})
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (db *Database) SpiderWebpageInsertOrSetUserCollected(title string, cover string, description string, metaDescription string, fullText string, url string, urlHash string, siteDomain string, userCollected int64) error {
	opts := options.FindOneAndUpdateOptions{}
	opts.SetUpsert(true)
	opts.SetProjection(bson.D{{"_id", 1}})

	var webpage model.Webpage

	now := time.Now()
	err := db.Database("note").Collection("webpage").FindOneAndUpdate(context.TODO(), bson.M{"url_hash": urlHash}, bson.M{"$setOnInsert": bson.M{"title": title, "cover": cover, "description": description, "meta_description": metaDescription, "full_text": fullText, "url": url, "url_hash": urlHash, "site_domain": siteDomain, "created_at": now}, "$set": bson.M{"user_collected": userCollected, "updated_at": now}}, &opts).Decode(&webpage)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	return nil
}

func (db *Database) SpiderCollectCreate(userId *primitive.ObjectID, urlHash string, cover string, description string, metaDescription string, fullText string, userCollected int64) (bool, error) {
	opts := options.FindOneAndUpdateOptions{}
	opts.SetProjection(bson.D{{"_id", 1}})

	var collect model.UserIdTagIdsCollect

	now := time.Now()
	err := db.Database("note").Collection("userid_tagids_collect_map").FindOneAndUpdate(context.TODO(), bson.M{"user_id": userId, "url_hash": urlHash, "crawl_status": "pending"}, bson.M{"$set": bson.M{"cover": cover, "description": description, "meta_description": metaDescription, "full_text": fullText, "user_collected": userCollected, "created_at": now, "updated_at": now, "crawl_status": "finished"}}, &opts).Decode(&collect)
	if err == mongo.ErrNoDocuments {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (db *Database) SpiderCollectUpdateStatus(userId *primitive.ObjectID, urlHash string, description string, crawlStatus string) (bool, error) {
	opts := options.FindOneAndUpdateOptions{}
	opts.SetProjection(bson.D{{"_id", 1}})

	var collect model.UserIdTagIdsCollect

	now := time.Now()
	var update primitive.M

	if crawlStatus == "failed" || crawlStatus == "unsupported" {
		update = bson.M{"$set": bson.M{"crawl_status": crawlStatus, "description": description, "updated_at": now}}
	} else {
		update = bson.M{"$set": bson.M{"crawl_status": crawlStatus, "updated_at": now}}
	}

	err := db.Database("note").Collection("userid_tagids_collect_map").FindOneAndUpdate(context.TODO(), bson.M{"user_id": userId, "url_hash": urlHash}, update, &opts).Decode(&collect)
	if err == mongo.ErrNoDocuments {
		return false, err
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
