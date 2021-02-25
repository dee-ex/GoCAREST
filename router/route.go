package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HandlerFunc is normal handler
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Route is a struct contains route information
type Route struct {
	Path        string
	Alias       string
	Name        string
	Methods     []string
	Handler     HandlerFunc
	Middlewares []mux.MiddlewareFunc
}

// NewRoute creates a route for router
func NewRoute(path, alias, name string, methods []string, handler HandlerFunc, middlewares ...mux.MiddlewareFunc) *Route {
	r := &Route{}

	r.Path = path
	r.Alias = alias
	r.Name = name
	r.Methods = methods
	r.Handler = handler
	r.Middlewares = middlewares

	return r
}
