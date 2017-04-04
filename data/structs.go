package data

import (
	"gopkg.in/mgo.v2/bson"
)

type RawPhoto struct {
	ID    bson.ObjectId `json:"id" bson:"_id"`
	Owner bson.ObjectId `json:"owner" bson:"owner"`
	URL   string        `json:"url"`
}

type User struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
}
