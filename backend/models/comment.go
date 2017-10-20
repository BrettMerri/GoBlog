package models

import "gopkg.in/mgo.v2/bson"

// Comment model
type Comment struct {
	ID      bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Body    string        `json:"body" bson:"body"`
	User    User          `json:"user" bson:"user"`
	Article bson.ObjectId `json:"article,omitempty" bson:"-"`
}
