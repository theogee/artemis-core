package main

import (
	"github.com/julienschmidt/httprouter"
	artemis "github.com/theogee/artemis-core/internal/handler"
	"github.com/theogee/artemis-core/pkg/config"
)

func registerRoutes(cfg *config.Config, router *httprouter.Router, artemisHandler *artemis.ArtemisHandler) {
	router.POST("/api/adm/register", artemisHandler.RegisterAsAdmin)
	router.POST("/api/adm/login", artemisHandler.LoginAsAdmin)

	// protected routes
	router.POST("/api/adm/logout", artemisHandler.Authenticate(artemisHandler.Logout, "adm"))
}
