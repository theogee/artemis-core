package main

import (
	"github.com/theogee/artemis-core/pkg/config"
	"github.com/theogee/artemis-core/pkg/connection/cache"
	"github.com/theogee/artemis-core/pkg/connection/database"
	"github.com/theogee/artemis-core/pkg/logger"
)

func main() {
	logger.Setup("main.log")
	defer logger.Close()

	var (
		logPrefix = "[main]"
		log       = logger.Log
	)

	log.Printf("%v hello from artemis-core!", logPrefix)

	cfg := config.Load("resources/config/artemis-core.json")

	c := cache.NewCache(cfg)

	db := database.NewDatabase(cfg)

	startApp(cfg, db, c)
}
