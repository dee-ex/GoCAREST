package router

import (
	"net/http"

	"github.com/rs/cors"
)

// NewCors creates cors for router
func NewCors() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
	})

	return c
}
