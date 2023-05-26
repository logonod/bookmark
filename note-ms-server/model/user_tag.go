package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserIdTag struct {
	Model `bson:",inline"`

	User *primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Name *string             `bson:"tag_name" json:"name,omitempty"`
}
