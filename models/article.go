package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionArticle holds the name of the articles collection
	CollectionArticle = "articles"
)

// Article model
type Article struct {
	ID    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string        `json:"title" bson:"title"`
	Body  string        `json:"body" bson:"body"`
	// User      bson.ObjectId `json:"user"`
}
