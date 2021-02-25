package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// SubRter is a struct contains subrouter information
// This struct helps us keep track when init subrouter
type SubRter struct {
	Object        *mux.Router
	HasMiddleware bool
}

// NewRouter creates a router for our application
func NewRouter() *mux.Router {
	rr := Routes()

	rter := mux.NewRouter()

	subRters := make(map[string]*SubRter)

	for _, r := range rr {
		if r.Name != "root" {
			if subRters[r.Name] == nil {
				subRter := &SubRter{}
				subRter.Object = rter.PathPrefix("").Subrouter()

				subRters[r.Name] = subRter
			}
		}
	}

	for _, r := range rr {
		if r.Name == "root" {
			rter.HandleFunc(r.Path, r.Handler).Methods(r.Methods...)
		} else {
			subRter := subRters[r.Name]
			if !subRter.HasMiddleware {
				for _, mdware := range r.Middlewares {
					subRter.Object.Use(mdware)
				}
				subRter.HasMiddleware = true
			}
			subRter.Object.HandleFunc(r.Path, r.Handler).Methods(r.Methods...)
		}
	}

	rter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Not found",
		})
	})

	return rter
}
