package routing

import (
	"net/http"

	"github.com/carlso70/go-todo/handlers"
)

// Route contains the name, method(GET,POST), pattern(URL),
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes lots of route
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		handlers.TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		handlers.TodoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todoCreate",
		handlers.TodoCreate,
	},
}
