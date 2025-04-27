package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route is a struct for all routes in the API
type Route struct {
	URI                   string
	Method                string
	Function              func(w http.ResponseWriter, r *http.Request)
	RequireAuthentication bool
}

// Config is a function that get a router and return it with all routes configured
func Config(r *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
