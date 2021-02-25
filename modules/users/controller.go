package users

import (
	"encoding/json"
	"net/http"
)

type cont struct {
	serv Service
}

// NewController creates a controller coresspond with our Service
func NewController(serv Service) *cont {
	return &cont{serv: serv}
}

// Handler is function takes cares request
func Handler(w http.ResponseWriter, r *http.Request) {
	mysqlRepo, err := NewMySQLRepository()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	userServ, err := NewService(mysqlRepo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	cont := NewController(userServ)

	u, err := cont.serv.Find(0)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}
