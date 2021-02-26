package router

import (
	"net/http"

	"github.com/dee-ex/gocarest/infra"
	"github.com/dee-ex/gocarest/mdwares"
	"github.com/dee-ex/gocarest/modules/users"
)

// Routes creates routes for our application
// Implement your routes here
func Routes() []*Route {
	db := infra.DBInitialization()
	repo := users.NewRepository(db)
	serv := users.NewService(repo)
	cont := users.NewController(serv)

	rr := []*Route{
		NewRoute("/", "helloworld", "root", []string{http.MethodGet}, cont.Handler),

		NewRoute("/me", "me", "auth", []string{http.MethodGet}, cont.Handler, mdwares.AuthMiddleware),
	}

	return rr
}
