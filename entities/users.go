package entities

import (
	"time"
)

// User is a struct contains user information
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName configs table of this entity in our database
func (User) TableName() string {
	return "users"
}
