package router

import (
	"net/http"
	"os"

	"github.com/dee-ex/gocarest/infra"
	"github.com/dee-ex/gocarest/mdwares"
	"github.com/dee-ex/gocarest/modules/users"
)

// Routes creates routes for our application
// Implement your routes here
func Routes() []*Route {
	db := infra.DBInitialization()
	cache := users.NewCache(os.Getenv("CACHE_HOST")+":"+os.Getenv("CACHE_PORT"), os.Getenv("CACHE_PASS"), 0, 10)
	repo := users.NewRepository(db, cache)
	serv := users.NewService(repo)
	cont := users.NewController(serv)

	rr := []*Route{
		NewRoute("/", "helloworld", "root", []string{http.MethodGet}, cont.HelloWorld),

		NewRoute("/users", "find", "root", []string{http.MethodGet}, cont.GetHandler),
		NewRoute("/users/{id}", "find-detail", "root", []string{http.MethodGet}, cont.GetDetailHandler),
		NewRoute("/users", "store", "root", []string{http.MethodPost}, cont.PostHandler),

		NewRoute("/me", "me", "auth", []string{http.MethodGet}, cont.HelloWorld, mdwares.AuthMiddleware),
	}

	return rr
}
