package cache

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/theogee/artemis-core/pkg/config"
	"github.com/theogee/artemis-core/pkg/logger"
)

type (
	Cache struct {
		Client *redis.Client
		Config *config.Config
	}
)

func NewCache(cfg *config.Config) *Cache {
	var (
		logPrefix = "[cache.NewCache]"
		log       = logger.Log

		c = &Cache{
			Config: cfg,
		}
	)

	err := c.connect()
	if err != nil {
		log.Fatalf("%v error connecting to artemis-core-redis. err: %v", logPrefix, err)
	}

	return c
}

func (c *Cache) connect() error {
	var (
		logPrefix = "[cache.connect]"
		log       = logger.Log

		address     = fmt.Sprintf("%v:%v", c.Config.Cache.Host, c.Config.Cache.Port)
		redisClient = redis.NewClient(&redis.Options{
			Addr:     address,
			Password: c.Config.Cache.Password,
			DB:       c.Config.Cache.DB,
		})
	)

	status := redisClient.Ping()
	if status.Err() != nil {
		log.Printf("%v error pinging artemis-core-redis. err: %v", logPrefix, status.Err())
		return status.Err()
	}

	log.Printf("%v connected to artemis-core-redis", logPrefix)
	c.Client = redisClient

	return nil
}
