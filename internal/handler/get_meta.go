package artemis

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
	utilsHTTP "github.com/theogee/artemis-core/pkg/utils/http"
)

func (h *ArtemisHandler) GetMeta(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		logPrefix = "[artemis.ArtemisHandler.GetMeta]"
		log       = logger.Log

		statusCode = http.StatusInternalServerError

		resp = &model.DefaultResponse{
			Success:   false,
			ServError: []string{},
		}

		d = &model.GetMetaResponse{}

		cookies = []*http.Cookie{}
	)

	defer func() {
		resp.Data = d
		utilsHTTP.SendJSON(w, cookies, resp, statusCode)
	}()

	sid := ps.ByName("sid")

	c, err := h.artemisUsecase.GetUserCache(sid)
	if err != nil {
		log.Printf("%v error calling artemisUsecase.GetUserCache. err: %v", logPrefix, err)
		resp.ServError = append(resp.ServError, err.Error())
		return
	}

	d.UserType = c.UserType
	resp.Success = true
	statusCode = http.StatusOK
}
