package data

import (
	"net/http"

	"github.com/graphql-go/handler"
)

type RelayMiddleware struct {
	h *handler.Handler
}

func CreateRelayMiddleware() *RelayMiddleware {
	// simplest relay-compliant graphql server HTTP handler
	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	return &RelayMiddleware{h}
}

func (rm *RelayMiddleware) AttachMiddleware() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rm.h.ContextHandler(r.Context(), w, r)
	}
}
