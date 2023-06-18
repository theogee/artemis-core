package artemis

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
	utilHTTP "github.com/theogee/artemis-core/pkg/utils/http"
)

func (h *ArtemisHandler) Logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		logPrefix = "[artemis.ArtemisHandler.Logout]"
		log       = logger.Log

		statusCode = http.StatusInternalServerError

		resp = &model.DefaultResponse{
			Success:   false,
			ServError: []string{},
		}

		d = &model.LogoutResponse{
			ErrMessage: []string{},
			Message:    []string{},
		}

		cookies = []*http.Cookie{}
	)

	defer func() {
		resp.Data = d
		utilHTTP.SendJSON(w, cookies, resp, statusCode)
	}()

	// assumption: sid is guaranteed filled, otherwise the request won't go through the Authenticate middleware
	sid := ps.ByName("sid")

	err := h.artemisUsecase.Logout(sid)
	if err != nil {
		log.Printf("%v error calling artemisUsecase.Logout. err: %v", logPrefix, err)
		resp.ServError = append(resp.ServError, err.Error())
		return
	}

	cookie := &http.Cookie{
		Name:   ps.ByName("cookieName"),
		MaxAge: -1,
	}

	cookies = append(cookies, cookie)

	resp.Success = true
	statusCode = http.StatusOK
}
