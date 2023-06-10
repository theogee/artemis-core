package main

import (
	"github.com/theogee/artemis-core/pkg/config"
	"github.com/theogee/artemis-core/pkg/connection/cache"
	"github.com/theogee/artemis-core/pkg/connection/database"
	"github.com/theogee/artemis-core/pkg/logger"
)

func main() {
	logger.Setup("main.log")

	var (
		logPrefix = "[main]"
		log       = logger.Log
	)

	cfg := config.Load("resources/config/artemis-core.json")

	cache.NewCache(cfg)

	database.NewDatabase(cfg)

	log.Printf("%v hello from artemis-core!", logPrefix)

	logger.Close()
}
