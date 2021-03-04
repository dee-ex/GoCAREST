package users

import (
	"fmt"
	"log"
	"time"

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
	HelloWorld() (map[string]string, error)
}

type repo struct {
	db    *gorm.DB
	cache UserCache
}

// NewRepository creates a repository coresspond with a gorm db
func NewRepository(db *gorm.DB, cache UserCache) *repo {
	return &repo{db: db, cache: cache}
}

func (r *repo) HelloWorld() (map[string]string, error) {
	return map[string]string{
		"message": "Hello, world!",
	}, nil
}

func (r *repo) Find(id int) (*entities.User, error) {
	uCache, err := r.cache.Get(fmt.Sprint(id))
	if err == nil {
		return uCache, nil
	} else if err.Error() != "redis: nil" {
		log.Println(fmt.Sprintf("{\"error\": \"%s\", \"time\": \"%s\", \"api\": \"/users\"}", err.Error(), time.Now()))
	}

	var u entities.User
	res := r.db.Where("id = ?", id).Find(&u)

	r.cache.Set(fmt.Sprint(u.ID), &u)

	return &u, res.Error
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
	var uu []*entities.User
	res := r.db.Offset(offset).Limit(limit).Find(&uu)

	for _, u := range uu {
		r.cache.Set(fmt.Sprint(u.ID), u)
	}

	return uu, res.Error
}

func (r *repo) Store(u *entities.User) error {
	res := r.db.Create(u)

	r.cache.Set(fmt.Sprint(u.ID), u)
	fmt.Println("set")

	return res.Error
}

func (r *repo) Update(u *entities.User) error {
	return nil
}

func (r *repo) Delete(u *entities.User) error {
	return nil
}
