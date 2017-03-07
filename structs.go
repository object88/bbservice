package bbservice

import "gopkg.in/mgo.v2/bson"

type RawPhoto struct {
	ID    bson.ObjectId `json:"id" bson:"_id"`
	Owner bson.ObjectId `json:"owner" bson:"owner"`
}

type User struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
}
