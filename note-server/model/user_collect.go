package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserIdTagIdsCollect struct {
	Model `bson:",inline" json:",inline"`

	Type          *string               `bson:"type,omitempty" json:"type,omitempty"`
	User          *primitive.ObjectID   `bson:"user_id,omitempty" json:"-"`
	Tags          []*primitive.ObjectID `bson:"tag_ids,omitempty" json:"-"`
	Title         *string               `bson:"title,omitempty" json:"title,omitempty"`
	Cover         *string               `bson:"cover,omitempty" json:"cover,omitempty"`
	Description   *string               `bson:"description,omitempty" json:"description,omitempty"`
	Url           *string               `bson:"url,omitempty" json:"url,omitempty"`
	UrlHash       *string               `bson:"url_hash,omitempty" json:"url_hash,omitempty"`
	SiteDomain    *string               `bson:"site_domain,omitempty" json:"site_domain"`
	CrawlStatus   *string               `bson:"crawl_status,omitempty" json:"-"`
	UserCollected *int                  `bson:"user_collected,omitempty" json:"user_collected,omitempty"`
}

type UserIdUrlHash struct {
	User      *primitive.ObjectID `json:"user_id"`
	UrlHash   *string             `json:"url_hash"`
	CreatedAt *time.Time          `json:"created_at"`
}

func (m *UserIdUrlHash) Encode() ([]byte, error) {
	encoded, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}
