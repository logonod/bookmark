package model

import (
  "crypto/rand"
  "time"

  "go.mongodb.org/mongo-driver/bson/primitive"
)

type Id []byte

func NewId() Id {
  ret := make(Id, 20)
  if _, err := rand.Read(ret); err != nil {
    panic(err)
  }
  return ret
}

type Model struct {
  ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
  CreatedAt *time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
  UpdatedAt *time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
