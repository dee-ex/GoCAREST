package users

import (
	"gorm.io/gorm"

	"github.com/dee-ex/gocarest/entities"
	"github.com/dee-ex/gocarest/infra"
)

type repo struct {
	db *gorm.DB
}

// NewMySQLRepository creates a repository coresspond with a MySQL database
func NewMySQLRepository() (*repo, error) {
	db, err := infra.NewMySQLSession()

	return &repo{db: db}, err
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
