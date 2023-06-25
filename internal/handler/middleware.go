package artemis

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
	utilHTTP "github.com/theogee/artemis-core/pkg/utils/http"
)

func (h *ArtemisHandler) Authenticate(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var (
			logPrefix = "[artemis.ArtemisHandler.AuthenticateAdmin]"
			log       = logger.Log

			statusCode = http.StatusUnauthorized

			resp = &model.DefaultResponse{
				Success:   false,
				ServError: []string{},
			}

			d = &model.AuthenticateResponse{
				ErrMessage: []string{},
				Message:    []string{},
			}

			cookies = []*http.Cookie{}

			uid        string
			cookieName = h.cfg.API.AuthCookieName
		)

		defer func() {
			if uid == "" {
				resp.Data = d
				utilHTTP.SendJSON(w, cookies, resp, statusCode)
				return
			}

			n(w, r, ps)
		}()

		cookie, err := r.Cookie(cookieName)
		if err != nil {
			log.Printf("%v error cookie: %v can't be found. err: %v", logPrefix, cookieName, err)
			d.ErrMessage = append(d.ErrMessage, model.UnauthorizedAccess)
			return
		}

		sid := cookie.Value

		uid, err = h.artemisUsecase.Authenticate(sid)
		if err != nil {
			log.Printf("%v error calling artemisUsecase.Authenticate. err: %v", logPrefix, err)
			resp.ServError = append(resp.ServError, err.Error())
			statusCode = http.StatusInternalServerError
			return
		}

		ps = append(ps, httprouter.Param{Key: "uid", Value: uid}, httprouter.Param{Key: "sid", Value: sid})
	}
}
