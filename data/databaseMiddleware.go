package data

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"sync"

	"gopkg.in/mgo.v2"
)

type key struct {
	id int64
}

var sessionKey key

var once sync.Once
var session *DatabaseSession

type DatabaseSession struct {
	session *mgo.Session
}

func CreateDatabaseSession() *DatabaseSession {
	once.Do(func() {
		sessionKey = key{id: rand.Int63()}

		s, err := mgo.Dial("127.0.0.1:27017/")

		if err != nil {
			log.Printf("Failed to create session to database: %s\n", err)
			panic(err)
		}

		log.Printf("Connected to mongo database.\n")
		session = &DatabaseSession{session: s}
	})
	return session
}

// GetSessionFromContext retreives the one and only MongoDB session
// from the HTTP context
func GetSessionFromContext(context context.Context) *mgo.Session {
	return context.Value(sessionKey).(*mgo.Session)
}

// AttachMiddleware is a piece of middleware that clones the database
// session and puts in on the HTTP request context.
func (ds *DatabaseSession) AttachMiddleware() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			connection := session.session.Clone()

			ctx := context.WithValue(r.Context(), sessionKey, connection)

			go func() {
				select {
				case <-ctx.Done():
					connection.Close()
				}
			}()

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
