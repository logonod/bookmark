package model

import ()

type Webpage struct {
	Model `bson:",inline"`

	Title           *string `bson:"title,omitempty" json:"title,omitempty"`
	Cover           *string `bson:"cover,omitempty" json:"cover,omitempty"`
	Description     *string `bson:"description,omitempty" json:"description,omitempty"`
	MetaDescription *string `bson:"meta_description,omitempty" json:"meta_description,omitempty"`
	FullText        *string `bson:"full_text,omitempty" json:"full_text,omitempty"`
	Url             *string `bson:"url,omitempty" json:"url,omitempty"`
	UrlHash         *string `bson:"url_hash,omitempty" json:"url_hash,omitempty"`
	SiteDomain      *string `bson:"site_domain,omitempty" json:"site_domain,omitempty"`
	UserCollected   *int    `bson:"user_collected,omitempty" json:"user_collected,omitempty"`
}
