package main

import (
	"log"
	"net/http"

	"github.com/object88/bbservice/data"
)

func main() {
	ds := data.CreateDatabaseSession()
	r := data.CreateRelayMiddleware()

	// create graphql endpoint
	http.Handle("/graphql", data.Chain(r.AttachMiddleware(), ds.AttachMiddleware(), data.CorsMiddleware()))
	http.HandleFunc("/image", data.Chain(data.StreamImage, ds.AttachMiddleware()))

	// serve!
	port := ":8081"
	log.Printf(`GraphQL server starting up on https://localhost%s`, port)
	err := http.ListenAndServeTLS(port, "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatalf("ListenAndServe failed, %v", err)
	}
}
