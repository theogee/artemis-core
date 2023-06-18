package utils

import (
	"encoding/json"
	"net/http"

	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
)

func SendJSON(w http.ResponseWriter, cookies []*http.Cookie, data interface{}, statusCode int) {
	var (
		logPrefix = "[utils_http.SendJSON]"
		log       = logger.Log

		d []byte
	)

	defer func() {
		w.Header().Set("content-type", "application/json")

		for _, c := range cookies {
			http.SetCookie(w, c)
		}

		w.WriteHeader(statusCode)

		w.Write(d)
	}()

	d, err := json.Marshal(data)
	if err != nil {
		log.Printf("%v error marshalling response data. err: %v", logPrefix, err)

		statusCode = http.StatusInternalServerError

		errResponse := &model.DefaultResponse{
			Success:   false,
			ServError: []string{err.Error()},
		}

		d, _ = json.Marshal(errResponse)
	}
}
