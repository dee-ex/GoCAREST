package users

import (
	"github.com/dee-ex/gocarest/entities"
)

// Reader is an interface for retrieving data from database
type Reader interface {
	Find(id int) (*entities.User, error)
	FindByUsername(usrnm string) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	FindByToken(token string) (*entities.User, error)
	FindAll(offset, limit int) ([]*entities.User, error)
}

// Writer is interface for storing, editing & deleting data in database
type Writer interface {
	Store(u *entities.User) error
	Update(u *entities.User) error
	Delete(u *entities.User) error
}

// Repository an interface abstracts our database implementation
type Repository interface {
	Reader
	Writer
}
