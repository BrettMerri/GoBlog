package models

import "gopkg.in/mgo.v2/bson"

// Comment model
type Comment struct {
	ID      int           `json:"id,omitempty" bson:"id,omitempty"`
	Body    string        `json:"body" bson:"body"`
	User    bson.ObjectId `json:"user" bson:"user"`
	Article bson.ObjectId `json:"article,omitempty" bson:"-"`
}
