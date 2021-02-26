package users

import (
	"gorm.io/gorm"

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

// Repository is an interface abstracts our database implementation
type Repository interface {
	Reader
	Writer
}

type repo struct {
	db *gorm.DB
}

// NewRepository creates a repository coresspond with a gorm db
func NewRepository(db *gorm.DB) *repo {
	return &repo{db: db}
}

func (r *repo) Find(id int) (*entities.User, error) {
	return &entities.User{Username: "Hello, world!"}, nil
}

func (r *repo) FindByUsername(usrnm string) (*entities.User, error) {
	return nil, nil
}

func (r *repo) FindByEmail(email string) (*entities.User, error) {
	return nil, nil
}

func (r *repo) FindByToken(token string) (*entities.User, error) {
	return nil, nil
}

func (r *repo) FindAll(offset, limit int) ([]*entities.User, error) {
	return nil, nil
}

func (r *repo) Store(u *entities.User) error {
	return nil
}

func (r *repo) Update(u *entities.User) error {
	return nil
}

func (r *repo) Delete(u *entities.User) error {
	return nil
}
