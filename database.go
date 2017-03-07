package bbservice

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getCollection(s *mgo.Session, dbName string, collectionName string) *mgo.Collection {
	collection := s.DB(dbName).C(collectionName)
	return collection
}

// GetRawPhotos returns all photos
func GetRawPhotos() []*RawPhoto {
	s := GetSession()
	c := getCollection(s, "bbgraph", "rawphotos")
	query := c.Find(nil)
	var photos []*RawPhoto
	query.All(&photos)
	if photos != nil {
		return photos
	}
	return []*RawPhoto{}
}

// GetRawPhoto returns a *RawPhoto with the matching id
func GetRawPhoto(id string) *RawPhoto {
	oid := bson.ObjectIdHex(id)
	s := GetSession()
	c := getCollection(s, "bbgraph", "rawphotos")
	query := c.Find(bson.M{"_id": oid})
	var u *RawPhoto
	query.One(u)
	return u
}

// GetUser returns a *User with the matching id
func GetUser(id string) *User {
	oid := bson.ObjectIdHex(id)
	s := GetSession()
	c := getCollection(s, "bbgraph", "users")
	query := c.Find(bson.M{"_id": oid})
	var u *User
	query.One(u)
	return u
}

// GetViewer returns the current user
func GetViewer() *User {
	return GetUser("000000000000000000000001")
}
