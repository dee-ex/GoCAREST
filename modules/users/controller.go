package users

import (
	"encoding/json"
	"net/http"
)

// Controller is an interface abstracts our requests handling
type Controller interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type cont struct {
	serv Service
}

// NewController creates a controller coresspond with our Service
func NewController(serv Service) *cont {
	return &cont{serv: serv}
}

// Handler is function takes cares request
func (c *cont) Handler(w http.ResponseWriter, r *http.Request) {
	u, err := c.serv.Find(0)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}
