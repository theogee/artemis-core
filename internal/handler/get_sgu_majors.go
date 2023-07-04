package artemis

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
	utilsHTTP "github.com/theogee/artemis-core/pkg/utils/http"
)

func (h *ArtemisHandler) GetSGUMajors(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		logPrefix = "[artemis.ArtemisHandler.GetSGUMajors]"
		log       = logger.Log

		statusCode = http.StatusInternalServerError

		resp = &model.DefaultResponse{
			Success:   false,
			ServError: []string{},
		}

		d = &model.GetSGUMajorsResponse{
			Majors: []*model.SGUMajor{},
		}

		cookies = []*http.Cookie{}
	)

	defer func() {
		resp.Data = d
		utilsHTTP.SendJSON(w, cookies, resp, statusCode)
	}()

	SGUMajors, err := h.artemisUsecase.GetSGUMajors()
	if err != nil {
		log.Printf("%v error calling artemisUsecase.GetSGUMajors. err: %v", logPrefix, err)
		resp.ServError = append(resp.ServError, err.Error())
		return
	}

	statusCode = http.StatusOK
	resp.Success = true
	d.Majors = SGUMajors
}
