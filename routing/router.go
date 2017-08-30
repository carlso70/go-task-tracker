package routing

import (
	"net/http"

	"github.com/carlso70/go-todo/handlers"
	"github.com/gorilla/mux"
)

// NewRouter - Creates a new router and wraps a logger
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = handlers.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
