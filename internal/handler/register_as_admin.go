package artemis

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
	utilsHTTP "github.com/theogee/artemis-core/pkg/utils/http"
)

func (h *ArtemisHandler) RegisterAsAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		logPrefix = "[artemis.ArtemisHandler.RegisterAsAdmin]"
		log       = logger.Log

		statusCode = http.StatusBadRequest

		resp = &model.DefaultResponse{
			Success:   false,
			ServError: []string{},
		}

		d = &model.RegisterAsAdminResponse{
			ErrMessage: []string{},
			Message:    []string{},
		}
	)

	defer func() {
		resp.Data = d
		utilsHTTP.SendJSON(w, resp, statusCode)
	}()

	username := r.FormValue("username")
	if username == "" {
		log.Printf("%v error username can't be empty", logPrefix)
		d.ErrMessage = append(d.ErrMessage, model.UsernameCantBeEmpty)
	}

	password := r.FormValue("password")
	if password == "" {
		log.Printf("%v error password can't be empty", logPrefix)
		d.ErrMessage = append(d.ErrMessage, model.PasswordCantBeEmpty)
	}

	email := r.FormValue("email")
	if email == "" {
		log.Printf("%v error email can't be empty", logPrefix)
		d.ErrMessage = append(d.ErrMessage, model.EmailCantBeEmpty)
	}

	if len(d.ErrMessage) != 0 {
		return
	}

	data := &model.RegisterAsAdminRequest{
		Username: username,
		Password: password,
		Email:    email,
	}

	err := h.artemisUsecase.RegisterAsAdmin(data)
	if err != nil {
		if err.Error() == model.UsernameAlreadyExist {
			log.Printf("%v error registering admin with username: %v. err: username already exist", logPrefix, username)
			d.ErrMessage = append(d.ErrMessage, model.UsernameAlreadyExist)
			return
		}

		// internal server error
		log.Printf("%v error calling artemisUsecase.RegisterAsAdmin. err: %v", logPrefix, err)
		statusCode = http.StatusInternalServerError
		resp.ServError = append(resp.ServError, err.Error())
		return
	}

	statusCode = http.StatusCreated
	resp.Success = true
	d.Message = append(d.Message, fmt.Sprintf(model.AdminCreatedSuccessfully, username))
}
