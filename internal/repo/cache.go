package artemis

import (
	"time"

	"github.com/theogee/artemis-core/pkg/logger"
)

func (r *ArtemisRepo) SetCache(key string, value interface{}, exp time.Duration) error {
	var (
		logPrefix = "[artemis.ArtemisRepo.SetCache]"
		log       = logger.Log
	)

	err := r.cache.Client.Set(key, value, exp).Err()
	if err != nil {
		log.Printf("%v error setting cache to redis. err: %v", logPrefix, err)
	}

	return err
}

func (r *ArtemisRepo) GetCache(key string) (string, error) {
	var (
		logPrefix = "[artemis.ArtemisRepo.GetCache]"
		log       = logger.Log
	)

	val, err := r.cache.Client.Get(key).Result()
	if err != nil {
		log.Printf("%v error getting cache from redis. key: %v. err: %v", logPrefix, key, err)
	}

	return val, err
}

func (r *ArtemisRepo) DeleteCache(key string) error {
	var (
		logPrefix = "[artemis.ArtemisRepo.GetCache]"
		log       = logger.Log
	)

	err := r.cache.Client.Del(key).Err()
	if err != nil {
		log.Printf("%v error deleting cache from redis. key: %v. err: %v", logPrefix, key, err)
	}

	return err
}
