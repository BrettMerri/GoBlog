package models

import "gopkg.in/mgo.v2/bson"

// Article model
type Article struct {
	ID       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string        `json:"title" bson:"title"`
	Body     string        `json:"body" bson:"body"`
	User     bson.ObjectId `json:"user" bson:"user"`
	Comments []Comment     `json:"comments,omitempty" bson:"comments,omitempty"`
}
