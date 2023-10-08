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
			logPrefix = "[artemis.ArtemisHandler.Authenticate]"
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

			c          *model.UserCache
			cookieName = h.cfg.API.AuthCookieName
		)

		defer func() {
			if c == nil {
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

		c, err = h.artemisUsecase.Authenticate(sid)
		if err != nil {
			log.Printf("%v error calling artemisUsecase.Authenticate. err: %v", logPrefix, err)
			resp.ServError = append(resp.ServError, err.Error())
			statusCode = http.StatusInternalServerError
			return
		}

		ps = append(ps, httprouter.Param{Key: "uid", Value: c.UID}, httprouter.Param{Key: "sid", Value: sid}, httprouter.Param{Key: "userType", Value: c.UserType})
	}
}

func (h *ArtemisHandler) Authorize(n httprouter.Handle, authorizedType string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var (
			logPrefix = "[artemis.ArtemisHandler.Authorize]"
			log       = logger.Log

			statusCode = http.StatusForbidden

			resp = &model.DefaultResponse{
				Success:   false,
				ServError: []string{},
			}

			d = &model.AuthorizeResponse{
				ErrMessage: []string{},
				Message:    []string{},
			}

			userType string

			cookies = []*http.Cookie{}
		)

		defer func() {
			if userType != authorizedType {
				resp.Data = d
				utilHTTP.SendJSON(w, cookies, resp, statusCode)
				return
			}

			n(w, r, ps)
		}()

		userType = ps.ByName("userType")
		if userType != authorizedType {
			log.Printf("%v unauthorized attempt to access resource. uid: %v. api: %v", logPrefix, ps.ByName("uid"), r.URL.Path)
			d.ErrMessage = append(d.ErrMessage, "error unauthorized access")
			return
		}
	}
}
