package data

import (
	"net/http"

	"github.com/rs/cors"
)

func CorsMiddleware() Middleware {
	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"access-control-allow-methods", "access-control-allow-origin", "content-type"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedOrigins:   []string{"https://localhost:3001"},
		Debug:            true,
	})

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			c.ServeHTTP(w, r, next)
		}
	}
}
