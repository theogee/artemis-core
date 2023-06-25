package artemis

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
	utilsHTTP "github.com/theogee/artemis-core/pkg/utils/http"
)

func (h *ArtemisHandler) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		logPrefix = "[artemis.ArtemisHandler.Login]"
		log       = logger.Log

		statusCode = http.StatusBadRequest

		resp = &model.DefaultResponse{
			Success:   false,
			ServError: []string{},
		}

		d = &model.LoginResponse{
			ErrMessage: []string{},
			Message:    []string{},
		}

		cookies = []*http.Cookie{}
	)

	defer func() {
		resp.Data = d
		utilsHTTP.SendJSON(w, cookies, resp, statusCode)
	}()

	username := r.FormValue("username")
	if username == "" {
		log.Printf("%v error username can't be empty", logPrefix)
		d.ErrMessage = append(d.ErrMessage, model.UsernameCantBeEmpty)
		d.UsernameError = model.UsernameCantBeEmpty
	}

	password := r.FormValue("password")
	if password == "" {
		log.Printf("%v error password can't be empty", logPrefix)
		d.ErrMessage = append(d.ErrMessage, model.PasswordCantBeEmpty)
		d.PasswordError = model.PasswordCantBeEmpty
	}

	if len(d.ErrMessage) != 0 {
		return
	}

	data := &model.LoginRequest{
		Username: username,
		Password: password,
	}

	sid, err := h.artemisUsecase.Login(data)
	if err != nil {
		if err.Error() == model.IncorrectCredential {
			log.Printf("%v error incorrect username or password", logPrefix)

			d.UsernameError = model.IncorrectCredential
			d.PasswordError = model.IncorrectCredential
			d.ErrMessage = append(d.ErrMessage, model.IncorrectCredential)
			statusCode = http.StatusUnauthorized
			return
		}

		log.Printf("%v error calling artemisUsecase.Login. err: %v", logPrefix, err)
		resp.ServError = append(resp.ServError, err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	cookie := &http.Cookie{
		Name:     h.cfg.API.AuthCookieName,
		MaxAge:   h.cfg.API.AuthSessionExpiration,
		HttpOnly: true,
		Value:    sid,
	}

	cookies = append(cookies, cookie)

	statusCode = http.StatusOK
	resp.Success = true
}
