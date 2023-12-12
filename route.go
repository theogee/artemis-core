package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/julienschmidt/httprouter"
	artemis "github.com/theogee/artemis-core/internal/handler"
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/config"
)

func registerRoutes(cfg *config.Config, router *httprouter.Router, artemisHandler *artemis.ArtemisHandler) {
	router.POST("/api/adm/register", artemisHandler.RegisterAsAdmin)
	router.POST("/api/login", artemisHandler.Login)

	// protected routes
	router.POST("/api/logout", artemisHandler.Authenticate(artemisHandler.Logout))
	router.POST("/api/meta", artemisHandler.Authenticate(artemisHandler.GetMeta))
	router.GET("/api/students/:studentID", artemisHandler.Authenticate(artemisHandler.GetStudentByID))

	// admin only
	router.GET("/api/students", artemisHandler.Authenticate(artemisHandler.Authorize(artemisHandler.GetStudents, model.UserTypeAdmin)))

	router.GET("/api/sgu_majors", artemisHandler.Authenticate(artemisHandler.Authorize(artemisHandler.GetSGUMajors, model.UserTypeAdmin)))

	router.GET("/api/exchange_year", artemisHandler.Authenticate(artemisHandler.Authorize(artemisHandler.GetExchangeYear, model.UserTypeAdmin)))

	router.POST("/api/register_student_by_csv", artemisHandler.Authenticate(artemisHandler.Authorize(artemisHandler.RegisterStudentByCSV, model.UserTypeAdmin)))
}

func registerWebPlatform(cfg *config.Config, router *httprouter.Router) {
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wd, _ := os.Getwd()
		absolutePath := filepath.Join(wd, cfg.Service.StaticPath)
		indexPath := filepath.Join(absolutePath, "index.html")
		fileServer := http.FileServer(http.Dir(absolutePath))

		if !strings.HasPrefix(r.URL.Path, "/api/") {
			if strings.Contains(r.URL.Path, ".") {
				fileServer.ServeHTTP(w, r)
				return
			}

			http.ServeFile(w, r, indexPath)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
	})
}
