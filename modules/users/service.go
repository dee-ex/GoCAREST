package users

import (
	"github.com/dee-ex/gocarest/entities"
)

// Service is an interface abstracts our logic implementation
type Service interface {
	Reader
	Writer
}

type serv struct {
	repo Repository
}

// NewService create a service corresponding with our Repository
func NewService(repo Repository) *serv {
	return &serv{repo: repo}
}

func (s *serv) Find(id int) (*entities.User, error) {
	return s.repo.Find(id)
}

func (s *serv) FindByUsername(usrnm string) (*entities.User, error) {
	return nil, nil
}

func (s *serv) FindByEmail(email string) (*entities.User, error) {
	return nil, nil
}

func (s *serv) FindByToken(token string) (*entities.User, error) {
	return nil, nil
}

func (s *serv) FindAll(offset, limit int) ([]*entities.User, error) {
	return nil, nil
}

func (s *serv) Store(u *entities.User) error {
	return nil
}

func (s *serv) Update(u *entities.User) error {
	return nil
}

func (s *serv) Delete(u *entities.User) error {
	return nil
}
