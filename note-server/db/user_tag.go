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

func (db *Database) UserTagList(user *model.User, start_id *primitive.ObjectID) ([]*model.UserIdTag, error) {
	var tags []*model.UserIdTag

	options := options.Find()
	options.SetSort(bson.D{{"_id", -1}})
	options.SetLimit(15)
	options.SetProjection(bson.D{{"tag_name", 1}})

	cursor, err := db.Database("note").Collection("userid_tag_map").Find(context.TODO(), bson.M{"user_id": user.ID, "_id": bson.M{"$lt": start_id}}, options)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {

		var elem model.UserIdTag
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}

		tags = append(tags, &elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if tags == nil {
		tags = make([]*model.UserIdTag, 0)
	}

	return tags, nil
}

func (db *Database) UserTagCreate(user *model.User, name string) (*model.UserIdTag, error) {
	var userTag model.UserIdTag

	userTag.User = user.ID
	userTag.Name = &name
	now := time.Now()
	userTag.CreatedAt = &now
	userTag.UpdatedAt = &now

	r, err := db.Database("note").Collection("userid_tag_map").InsertOne(context.TODO(), userTag)
	if err != nil {
		return nil, err
	}

	id := r.InsertedID.(primitive.ObjectID)
	tag := &model.UserIdTag{Model: model.Model{ID: &id}, Name: &name}

	return tag, nil
}

func (db *Database) GetUserTagIdFromTagName(user *model.User, name string) (*primitive.ObjectID, error) {
	var tag model.UserIdTag

	options := options.FindOne()
	options.SetProjection(bson.D{{"_id", 1}})

	err := db.Database("note").Collection("userid_tag_map").FindOne(context.TODO(), bson.M{"user_id": user.ID, "tag_name": name}, options).Decode(&tag)
	if err != nil {
		return nil, err
	}

	return tag.ID, nil
}

func (db *Database) GetUserTagIdFromTagId(user *model.User, id *primitive.ObjectID) (*primitive.ObjectID, error) {
	var tag model.UserIdTag

	options := options.FindOne()
	options.SetProjection(bson.D{{"_id", 1}})

	err := db.Database("note").Collection("userid_tag_map").FindOne(context.TODO(), bson.M{"user_id": user.ID, "_id": id}, options).Decode(&tag)
	if err != nil {
		return nil, err
	}

	return tag.ID, nil
}

func (db *Database) InsertAndFindTagIdsFromTagNames(user *model.User, tagNames []string) ([]*primitive.ObjectID, error) {
	var writes []mongo.WriteModel
	var tagIds []*primitive.ObjectID

	now := time.Now()

	// 执行插入操作，确保每一条都有
	for _, elem := range tagNames {
		model := mongo.NewReplaceOneModel().SetFilter(bson.M{"user_id": user.ID, "tag_name": elem}).SetReplacement(bson.M{"$setOnInsert": bson.M{"created_at": now, "updated_at": now, "tag_name": elem}}).SetUpsert(true)
		writes = append(writes, model)
	}

	_, err := db.Database("note").Collection("userid_tag_map").BulkWrite(context.TODO(), writes)
	if err != nil {
		return nil, err
	}

	// 执行查询操作，获取Tag列表
	options := options.Find()
	options.SetLimit(3)
	cursor, err := db.Database("note").Collection("userid_tag_map").Find(context.TODO(), bson.M{"user_id": user.ID, "tag_name": bson.M{"$in": tagNames}}, options)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {

		var elem model.UserIdTag
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}

		tagIds = append(tagIds, elem.ID)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if tagIds == nil {
		tagIds = make([]*primitive.ObjectID, 0)
	}

	return tagIds, nil
}

func (db *Database) UserTagCollectCount(user *model.User, tag *primitive.ObjectID) (*int64, error) {
	var filter primitive.M
	if tag.Hex() == "000000000000000000000001" {
		filter = bson.M{"user_id": user.ID, "tag_ids": []primitive.ObjectID{}}
	} else {
		filter = bson.M{"user_id": user.ID, "tag_ids": tag}
	}

	count, err := db.Database("note").Collection("userid_tagids_collect_map").CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func (db *Database) UserTagDelete(tag *primitive.ObjectID) (*mongo.DeleteResult, error) {
	r, err := db.Database("note").Collection("userid_tag_map").DeleteOne(context.TODO(), bson.M{"_id": tag})
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (db *Database) UserTagUpdate(user *model.User, inputFrom string, inputTo string) (*mongo.UpdateResult, error) {
	r, err := db.Database("note").Collection("userid_tag_map").UpdateOne(context.TODO(), bson.M{"user_id": user.ID, "tag_name": inputFrom}, bson.M{"$set": bson.M{"tag_name": inputTo}})
	if err != nil {
		return nil, err
	}

	return r, nil
}
