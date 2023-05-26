package cache

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"e.coding.net/logonod/note-server/model"
)

func (cache *Cache) SetUserCookie(uid string, user *model.User, ttl int) error {
	if err := cache.Set(uid, (*user.ID).Hex(), time.Duration(ttl)*time.Second).Err(); err != nil {
		return err
	}

	return nil
}

func (cache *Cache) GetUserCookie(cookie string) (*model.User, error) {
	s, err := cache.Get(cookie).Result()
	if err != nil {
		return nil, err
	}

	id, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		return nil, err
	}

	return &model.User{Model: model.Model{ID: &id}}, nil
}

func (cache *Cache) DeleteUserCookie(uid string) error {
	if err := cache.Del(uid).Err(); err != nil {
		return err
	}

	return nil
}
