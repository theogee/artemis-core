package artemis

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
	utilsHTTP "github.com/theogee/artemis-core/pkg/utils/http"
)

func (h *ArtemisHandler) GetStudents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		logPrefix = "[artemis.ArtemisHandler.GetStudents]"
		log       = logger.Log

		statusCode = http.StatusBadRequest

		resp = &model.DefaultResponse{
			Success:   false,
			ServError: []string{},
		}

		d = &model.GetStudentsResponse{
			ErrMessage: []string{},
			Students:   []*model.StudentSimple{},
		}

		cookies = []*http.Cookie{}

		data = &model.GetStudentsRequest{}
	)

	defer func() {
		resp.Data = d
		utilsHTTP.SendJSON(w, cookies, resp, statusCode)
	}()

	// only admin can fetch student data
	sid := ps.ByName("sid")

	c, err := h.artemisUsecase.GetUserCache(sid)
	if err != nil {
		log.Printf("%v error calling artemisUsecase.GetUserCache. err: %v", logPrefix, err)
		statusCode = http.StatusInternalServerError
		resp.ServError = append(resp.ServError, err.Error())
		return
	}

	if c.UserType != model.UserTypeAdmin {
		log.Printf("%v error received request from non-admin user. UID: %v", logPrefix, c.UID)
		statusCode = http.StatusForbidden
		d.ErrMessage = append(d.ErrMessage, model.ForbiddenAccess)
		return
	}

	q := r.URL.Query()

	limitStr := q.Get("limit")
	if limitStr != "" {
		limit, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			log.Printf("%v error limit must be a valid number. err: %v", logPrefix, err)
			d.ErrMessage = append(d.ErrMessage, "limit must be a valid number")
		} else {
			if limit <= 0 {
				log.Printf("%v error limit must be greater than 0", logPrefix)
				d.ErrMessage = append(d.ErrMessage, "limit must be greater than 0")
			}

			data.Limit = limit
		}
	} else {
		data.Limit = 20
	}

	pageStr := q.Get("page")
	if pageStr != "" {
		page, err := strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			log.Printf("%v error page must be a valid number. err: %v", logPrefix, err)
			d.ErrMessage = append(d.ErrMessage, "page must be a valid number")
		} else {
			if page <= 0 {
				log.Printf("%v error page must not be greater than 0", logPrefix)
				d.ErrMessage = append(d.ErrMessage, "page must be greater than 0")
			}

			data.Page = page
		}
	} else {
		data.Page = 1
	}

	if len(d.ErrMessage) != 0 {
		return
	}

	students, err := h.artemisUsecase.GetStudents(data)
	if err != nil {
		log.Printf("%v error calling artemisUsecase.GetStudents. err: %v", logPrefix, err)
		statusCode = http.StatusInternalServerError
		resp.ServError = append(resp.ServError, err.Error())
		return
	}

	d.Students = students
	statusCode = http.StatusOK
	resp.Success = true
}
