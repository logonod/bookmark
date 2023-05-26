package cache

import (
  "github.com/go-redis/redis"
  "github.com/pkg/errors"
)

type Cache struct {
  *redis.Client
}

func New(config *Config) (*Cache, error) {
  cache := redis.NewClient(&redis.Options{
    Addr:     config.Addr,
    Password: config.Password,
    DB:       config.DB,
  })

  _, err := cache.Ping().Result()

  if err != nil {
    return nil, errors.Wrap(err, "unable to connect to redis")
  }

  return &Cache{cache}, nil
}
