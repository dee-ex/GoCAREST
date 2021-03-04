package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Controller is an interface abstracts our requests handling
type Controller interface {
	HelloWorld(w http.ResponseWriter, r *http.Request)
	GetHandler(w http.ResponseWriter, r *http.Request)
	GetDetailHandler(w http.ResponseWriter, r *http.Request)
	PostHandler(w http.ResponseWriter, r *http.Request)
}

type cont struct {
	serv Service
}

// NewController creates a controller coresspond with our Service
func NewController(serv Service) *cont {
	return &cont{serv: serv}
}

// Handler is function takes cares request
func (c *cont) HelloWorld(w http.ResponseWriter, r *http.Request) {
	msg, err := c.serv.HelloWorld()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, http.StatusOK, msg)
}

func (c *cont) GetHandler(w http.ResponseWriter, r *http.Request) {
	oset, lim := GetNumericParameter(r, "offset", 0), GetNumericParameter(r, "limit", 100)
	uu, err := c.serv.FindAll(oset, lim)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, http.StatusOK, uu)
}

func (c *cont) GetDetailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	u, err := c.serv.Find(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if u.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	JSONResponse(w, http.StatusOK, u)
}

// PostHandler serve Get request
func (c *cont) PostHandler(w http.ResponseWriter, r *http.Request) {
	var data UserCreationPayload
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ok := CheckUCPayload(data); !ok {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	u := MakeUser(data)

	err = c.serv.Store(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, http.StatusOK, u)
}
