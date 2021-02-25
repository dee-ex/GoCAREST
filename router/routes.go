package router

import (
	"net/http"

	"github.com/dee-ex/gocarest/mdwares"
	"github.com/dee-ex/gocarest/modules/users"
)

// Routes creates routes for our application
// Implement your routes here
func Routes() []*Route {
	rr := []*Route{
		NewRoute("/", "helloworld", "root", []string{http.MethodGet}, users.Handler),

		NewRoute("/me", "me", "auth", []string{http.MethodGet}, users.Handler, mdwares.AuthMiddleware),
	}

	return rr
}
