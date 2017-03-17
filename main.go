package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/object88/bbservice/data"
	"github.com/rs/cors"
)

func main() {

	// simplest relay-compliant graphql server HTTP handler
	h := handler.New(&handler.Config{
		Schema: &data.Schema,
		Pretty: true,
	})

	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"access-control-allow-methods", "access-control-allow-origin", "content-type"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedOrigins:   []string{"https://localhost:3001"},
		Debug:            true,
	})

	// create graphql endpoint
	http.Handle("/graphql", c.Handler(h))

	// serve!
	port := ":8081"
	log.Printf(`GraphQL server starting up on https://localhost%s`, port)
	err := http.ListenAndServeTLS(port, "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatalf("ListenAndServe failed, %v", err)
	}
}
