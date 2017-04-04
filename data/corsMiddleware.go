package data

import (
	"net/http"

	"github.com/rs/cors"
)

type CorsMiddleware struct {
	c *cors.Cors
}

func CreateCorsMiddleware() *CorsMiddleware {
	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"access-control-allow-methods", "access-control-allow-origin", "content-type"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedOrigins:   []string{"https://localhost:3001"},
		Debug:            true,
	})

	return &CorsMiddleware{c}
}

func (c *CorsMiddleware) AttachMiddleware() Middleware {

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			c.c.ServeHTTP(w, r, next)
		}
	}
}
