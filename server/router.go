package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"bitdeal.nl/server/routes"
)

/*
NewRouter ...
*/
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Add GET requests
	for _, route := range routes.GetRoutes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// Add POST requests
	for _, route := range routes.PostRoutes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
