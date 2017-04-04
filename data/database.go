package data

import (
	"context"
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getCollection(s *mgo.Session, dbName string, collectionName string) *mgo.Collection {
	collection := s.DB(dbName).C(collectionName)
	return collection
}

// GetRawPhotos returns all photos
func GetRawPhotos(ctx context.Context) []*RawPhoto {
	// s := GetSessionFromContext()
	// c := getCollection(s, "bbgraph", "rawphotos")
	// query := c.Find(nil)
	// var photos []*RawPhoto
	// query.All(&photos)
	// if photos != nil {
	// 	return photos
	// }
	// return []*RawPhoto{}
	userID := bson.NewObjectId()

	photos := make([]*RawPhoto, 10)
	for i := 0; i < 10; i++ {
		id := bson.NewObjectId()
		photos[i] = &RawPhoto{
			ID:    id,
			Owner: userID,
			URL:   fmt.Sprintf("https://localhost:8081/image?id=%s", id.Hex()),
		}
	}
	return photos
}

// GetRawPhoto returns a *RawPhoto with the matching id
func GetRawPhoto(ctx context.Context, id string) *RawPhoto {
	oid := bson.ObjectIdHex(id)
	s := GetSessionFromContext(ctx)
	c := getCollection(s, "bbgraph", "rawphotos")
	query := c.Find(bson.M{"_id": oid})
	var u *RawPhoto
	query.One(u)
	u.URL = fmt.Sprintf("https://localhost:8081/image?id=%s", u.ID.Hex())
	return u
}

// GetUser returns a *User with the matching id
func GetUser(ctx context.Context, id string) *User {
	oid := bson.ObjectIdHex(id)
	s := GetSessionFromContext(ctx)
	c := getCollection(s, "bbgraph", "users")
	query := c.Find(bson.M{"_id": oid})
	var u *User
	query.One(u)
	return u
}

// GetViewer returns the current user
func GetViewer(ctx context.Context) *User {
	return GetUser(ctx, "000000000000000000000001")
}
