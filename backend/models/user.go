package models

import "gopkg.in/mgo.v2/bson"

// User model
type User struct {
	ID       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username"`
}
