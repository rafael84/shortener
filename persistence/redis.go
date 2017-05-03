package persistence

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

func NewRedis(address, password string, database int) *Redis {
	return &Redis{
		cli: redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       database,
		}),
	}
}

const (
	key      = "urls"
	keyCount = "count"
)

type Redis struct {
	cli *redis.Client
}

func (m *Redis) Set(alias, url string) error {
	b, err := m.cli.HSetNX(key, alias, url).Result()
	if err != nil {
		return err
	}
	if !b {
		return fmt.Errorf("could not set alias[%v] url[%v]", alias, url)
	}
	return nil
}

func (m *Redis) Get(alias string) (url string, found bool) {
	s, err := m.cli.HGet(key, alias).Result()
	if err != nil || s == "" {
		return "", false
	}
	return s, true
}

func (m *Redis) Count() int {
	s, err := m.cli.Get(keyCount).Result()
	if s != "" && err == nil {
		if parsed, err := strconv.Atoi(s); err == nil {
			return parsed
		}
	}
	return 0
}

func (m *Redis) Increment() error {
	return m.cli.Incr(keyCount).Err()
}
