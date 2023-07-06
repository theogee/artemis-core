package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	artemisHandler "github.com/theogee/artemis-core/internal/handler"
	artemisRepo "github.com/theogee/artemis-core/internal/repo"
	artemisUsecase "github.com/theogee/artemis-core/internal/usecase"
	"github.com/theogee/artemis-core/pkg/config"
	"github.com/theogee/artemis-core/pkg/connection/cache"
	"github.com/theogee/artemis-core/pkg/connection/database"
	"github.com/theogee/artemis-core/pkg/logger"
)

func startApp(cfg *config.Config, db *database.Database, c *cache.Cache) {
	var (
		logPrefix = "[main.startApp]"
		log       = logger.Log
	)

	router := httprouter.New()

	artemisRepo := artemisRepo.NewRepo(cfg, db, c)

	artemisUsecase := artemisUsecase.NewUsecase(cfg, artemisRepo)

	artemisHandler := artemisHandler.NewHandler(cfg, artemisUsecase)

	registerRoutes(cfg, router, artemisHandler)

	addr := fmt.Sprintf("%v:%v", cfg.Service.Host, cfg.Service.Port)

	log.Printf("%v starting HTTP server at %v", logPrefix, addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatalf("%v error starting HTTP server. err: %v", logPrefix, err)
	}
}
