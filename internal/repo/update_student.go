package artemis

import (
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
)

func (r *ArtemisRepo) UpdateStudentByID(data *model.UpdateStudentByIDRequest) error {
	var (
		logPrefix = "[artemis.ArtemisRepo.UpdateStudentByID]"
		log       = logger.Log
	)

	_, err := r.db.Conn.Exec(UpdateStudentByIDQuery, data.MobilePhone, data.MobilePhoneDE, data.PrivateEmail, data.CurrentAddress, data.CurrentPostcode, data.CurrentCity, data.CoName, data.InternshipCompany, data.InternshipStartDate, data.InternshipEndDate, data.InternshipCompanyAddress, data.InternshipCompanyPostcode, data.InternshipCompanyCity, data.InternshipSupervisorName, data.InternshipSupervisorEmail, data.InternshipSupervisorPhone, data.SGUEmail, data.FHEmail, data.IBAN, data.StudentID)
	if err != nil {
		log.Printf("%v error updating student data to database. err: %v", logPrefix, err)
		return err
	}

	log.Printf("%v studentID: %v has been updated in database", logPrefix, data.StudentID)

	return nil
}
