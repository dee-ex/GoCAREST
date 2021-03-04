package users

import (
	"context"
	"encoding/json"
	"time"

	"github.com/dee-ex/gocarest/entities"
	"github.com/go-redis/redis/v8"
)

// UserCache is an interface abstracts our cache for user entity implementation
type UserCache interface {
	Set(key string, value *entities.User) error
	Get(key string) (*entities.User, error)
}

type cache struct {
	host    string
	pass    string
	db      int
	expires time.Duration
	ctx     context.Context
}

// NewCache creates a cache
func NewCache(host, pass string, db int, exp time.Duration) *cache {
	return &cache{
		host:    host,
		pass:    pass,
		db:      db,
		expires: exp,
		ctx:     context.Background(),
	}
}

func (c *cache) GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.host,
		Password: c.pass,
		DB:       c.db,
	})
}

func (c *cache) Set(key string, val *entities.User) error {
	cli := c.GetClient()

	json, err := json.Marshal(val)
	if err != nil {
		return err
	}

	cli.Set(c.ctx, key, json, c.expires*time.Second)

	return nil
}

func (c *cache) Get(key string) (*entities.User, error) {
	cli := c.GetClient()

	val, err := cli.Get(c.ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var u entities.User
	err = json.Unmarshal([]byte(val), &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
