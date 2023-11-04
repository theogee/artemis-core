package artemis

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
	utilsHTTP "github.com/theogee/artemis-core/pkg/utils/http"
)

func (h *ArtemisHandler) GetStudentByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		logPrefix = "[artemis.ArtemisHandler.GetStudentByID]"
		log       = logger.Log

		statusCode = http.StatusBadRequest

		resp = &model.DefaultResponse{
			Success:   false,
			ServError: []string{},
		}

		d = &model.GetStudentByIDResponse{
			ErrMessage: []string{},
			Student:    nil,
		}

		cookies = []*http.Cookie{}

		studentIDStr string
	)

	defer func() {
		resp.Data = d
		utilsHTTP.SendJSON(w, cookies, resp, statusCode)
	}()

	userType := ps.ByName("userType")

	if userType == model.UserTypeStudent {
		studentIDStr = ps.ByName("uid")
	} else {
		studentIDStr = ps.ByName("studentID")
	}

	studentID, err := strconv.ParseInt(studentIDStr, 10, 64)
	if err != nil {
		log.Printf("%v error studentID must be a valid number. err: %v", logPrefix, err)
		d.ErrMessage = append(d.ErrMessage, "studentID must be a valid number")
		return
	}

	student, err := h.artemisUsecase.GetStudentByID(studentID)
	if err != nil {
		log.Printf("%v error calling artemisUsecase.GetStudentByID. err: %v", logPrefix, err)
		statusCode = http.StatusInternalServerError
		resp.ServError = append(resp.ServError, err.Error())
		return
	}

	if student == nil {
		statusCode = http.StatusNotFound
		d.ErrMessage = append(d.ErrMessage, "studentID doesn't exist")
		return
	}

	statusCode = http.StatusOK
	d.Student = student
	resp.Success = true
}
