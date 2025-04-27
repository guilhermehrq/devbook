package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router and returns it
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
