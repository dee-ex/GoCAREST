package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dee-ex/gocarest/entities"
)

// Implement your helping functions here

// CheckUCPayload will verify UserCreationPayload for regist new user
func CheckUCPayload(data UserCreationPayload) bool {
	if data.Username == "" {
		return false
	}
	if data.Email == "" {
		return false
	}
	if data.Password == "" {
		return false
	}

	return true
}

// MakeUser creates a user from data payload
func MakeUser(data UserCreationPayload) *entities.User {
	return &entities.User{Username: data.Username, Email: data.Email, Password: data.Password}
}

// GetParameter gets URL parameter from corresponding request
func GetParameter(r *http.Request, key, defl string) string {
	val := r.URL.Query()[key]

	if len(val) == 0 {
		return defl
	}

	if len(val[0]) == 0 {
		return defl
	}

	return val[0]
}

// GetNumericParameter gets numeric URL parameter from correspoding request
func GetNumericParameter(r *http.Request, key string, defl int) int {
	val, err := strconv.Atoi(GetParameter(r, key, ""))
	if err != nil {
		return defl
	}

	return val
}

// JSONResponse responds output as JSON format
func JSONResponse(w http.ResponseWriter, httpStat int, msg interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStat)
	json.NewEncoder(w).Encode(msg)
}
