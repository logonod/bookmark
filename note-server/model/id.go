package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IdsContainsIds(ids []*primitive.ObjectID, s []*primitive.ObjectID) bool {
	if len(s) == 0 && len(ids) != 0 {
		return false
	}

	m := make(map[string]bool)
	for _, v := range ids {
		m[v.Hex()] = true
	}
	for _, v := range s {
		if !m[v.Hex()] {
			return false
		}
	}

	return true
}
