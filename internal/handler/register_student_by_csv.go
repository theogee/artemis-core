package artemis

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
	utilsHTTP "github.com/theogee/artemis-core/pkg/utils/http"
)

func (h *ArtemisHandler) RegisterStudentByCSV(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		logPrefix = "[artemis.ArtemisHandler.RegisterStudentByCSV]"
		log       = logger.Log

		statusCode = http.StatusInternalServerError

		resp = &model.DefaultResponse{
			Success:   false,
			ServError: []string{},
		}

		d = &model.RegisterStudentByCSV{
			ErrMessage: []string{},
			Message:    []string{},
		}

		cookies = []*http.Cookie{}
	)

	defer func() {
		resp.Data = d
		utilsHTTP.SendJSON(w, cookies, resp, statusCode)
	}()

	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		log.Printf("%v error parsing form data. err: %v", logPrefix, err)
		resp.ServError = append(resp.ServError, err.Error())
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("%v error getting file from request. err: %v", logPrefix, err)
		d.ErrMessage = append(d.ErrMessage, "file not found")
		statusCode = http.StatusBadRequest
		return
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)
	if ext != ".csv" {
		log.Printf("%v error file is not a csv", logPrefix)
		d.ErrMessage = append(d.ErrMessage, "file is not a csv")
		statusCode = http.StatusBadRequest
		return
	}

	if header.Size > (10 << 20) {
		log.Printf("%v error file greater than 10MB", logPrefix)
		d.ErrMessage = append(d.ErrMessage, "file must not be greater than 10MB")
		statusCode = http.StatusBadRequest
		return
	}

	exchangeYearStr := r.FormValue("exchangeYear")
	exchangeYear, err := strconv.ParseInt(exchangeYearStr, 10, 64)
	if err != nil {
		log.Printf("%v error exchangeYear must be a valid number. err: %v", logPrefix, err)
		d.ErrMessage = append(d.ErrMessage, "exchangeYear must be a valid number")
		statusCode = http.StatusBadRequest
		return
	}

	if exchangeYear <= 2000 || exchangeYear > 2900 {
		log.Printf("%v error exchangeYear value is outside the boundary", logPrefix)
		d.ErrMessage = append(d.ErrMessage, "exchangeYear value is outside the boundary")
		statusCode = http.StatusBadRequest
		return
	}

	err = h.artemisUsecase.RegisterStudentByCSV(file, header, exchangeYear)
	if err != nil {
		log.Printf("%v error calling artemisUsecase.RegisterStudentByCSV. err: %v", logPrefix, err)
		resp.ServError = append(resp.ServError, err.Error())
		return
	}

	resp.Success = true
	statusCode = http.StatusCreated
}
