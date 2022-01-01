package entities

import (
	"time"
)

type Post struct {
	Id          string    `json:"id" bson:"_id"`
	Description string    `json:"description" bson:"description"`
	Store       Store     `json:"store" bson:"store"`
	Rating      int8      `json:"rating" bson:"rating"`
	HashTags    []string  `json:"hash_tags" bson:"hash_tags"`
	DateCreated time.Time `json:"date_created" bson:"date_created"`
	Foods       []Food    `json:"foods" bson:"foods"`
}
