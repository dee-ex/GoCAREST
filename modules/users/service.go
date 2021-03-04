package users

import (
	"github.com/dee-ex/gocarest/entities"
)

// Service is an interface abstracts our logic implementation
type Service interface {
	Reader
	Writer
	HelloWorld() (map[string]string, error)
}

type serv struct {
	repo Repository
}

// NewService creates a service corresponding with our Repository
func NewService(repo Repository) *serv {
	return &serv{repo: repo}
}

func (s *serv) HelloWorld() (map[string]string, error) {
	return s.repo.HelloWorld()
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
	return s.repo.FindAll(offset, limit)
}

func (s *serv) Store(u *entities.User) error {
	return s.repo.Store(u)
}

func (s *serv) Update(u *entities.User) error {
	return nil
}

func (s *serv) Delete(u *entities.User) error {
	return nil
}
