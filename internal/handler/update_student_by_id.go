package artemis

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
	utilsHTTP "github.com/theogee/artemis-core/pkg/utils/http"
)

func (h *ArtemisHandler) UpdateStudentByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		logPrefix = "[artemis.ArtemisHandler.UpdateStudentByID]"
		log       = logger.Log

		statusCode = http.StatusBadRequest

		resp = &model.DefaultResponse{
			Success:   false,
			ServError: []string{},
		}

		d = &model.UpdateStudentByIDResponse{
			ErrMessage: []string{},
			Message:    []string{},
		}

		cookies = []*http.Cookie{}

		studentIDStr string

		currentPostcode, internshipCompanyPostcode, internshipStartDate, internshipEndDate interface{}
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

	currentPostcode = r.FormValue("currentPostcode")
	if currentPostcode == "" {
		currentPostcode = nil
	}

	internshipCompanyPostcode = r.FormValue("internshipCompanyPostcode")
	if internshipCompanyPostcode == "" {
		internshipCompanyPostcode = nil
	}

	internshipStartDate = r.FormValue("internshipStartDate")
	if internshipStartDate == "" {
		internshipStartDate = nil
	}

	internshipEndDate = r.FormValue("internshipEndDate")
	if internshipEndDate == "" {
		internshipEndDate = nil
	}

	data := &model.UpdateStudentByIDRequest{
		StudentID:                 studentID,
		MobilePhone:               r.FormValue("mobilePhone"),
		MobilePhoneDE:             r.FormValue("mobilePhoneDE"),
		PrivateEmail:              r.FormValue("privateEmail"),
		CurrentAddress:            r.FormValue("currentAddress"),
		CurrentPostcode:           currentPostcode,
		CurrentCity:               r.FormValue("currentCity"),
		CoName:                    r.FormValue("coName"),
		InternshipCompany:         r.FormValue("internshipCompany"),
		InternshipStartDate:       internshipStartDate,
		InternshipEndDate:         internshipEndDate,
		InternshipCompanyAddress:  r.FormValue("internshipCompanyAddress"),
		InternshipCompanyPostcode: internshipCompanyPostcode,
		InternshipCompanyCity:     r.FormValue("internshipCompanyCity"),
		InternshipSupervisorName:  r.FormValue("internshipSupervisorName"),
		InternshipSupervisorEmail: r.FormValue("internshipSupervisorEmail"),
		InternshipSupervisorPhone: r.FormValue("internshipSupervisorPhone"),
		SGUEmail:                  r.FormValue("sguEmail"),
		FHEmail:                   r.FormValue("fhEmail"),
		IBAN:                      r.FormValue("iban"),
	}

	err = h.artemisUsecase.UpdateStudentByID(data)
	if err != nil {
		if err.Error() == model.StudentIDDoesNotExist {
			log.Printf("%v error updating student with ID: %v. err: studentID doesn't exist", logPrefix, studentID)
			d.ErrMessage = append(d.ErrMessage, fmt.Sprintf(model.StudentIDDoesNotExist, studentID))
			return
		}

		// internal server error
		log.Printf("%v error calling artemisUsecase.UpdateStudentByID. err: %v", logPrefix, err)
		statusCode = http.StatusInternalServerError
		resp.ServError = append(resp.ServError, err.Error())
		return
	}

	statusCode = http.StatusOK
	resp.Success = true
	d.Message = append(d.Message, fmt.Sprintf(model.StudentWithIDUpdatedSuccessfully, studentID))
}
