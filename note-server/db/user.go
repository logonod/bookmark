package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"e.coding.net/logonod/note-server/model"
)

func (db *Database) UserLogin(user *model.User) error {

	err := db.Database("note").Collection("user").FindOne(context.TODO(), bson.M{"phone": user.Phone}).Decode(user)
	if err != nil {
		return err
	}

	return nil
}
