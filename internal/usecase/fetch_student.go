package artemis

import (
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
)

func (u *ArtemisUsecase) GetStudents(data *model.GetStudentsRequest) ([]*model.StudentSimple, int, error) {
	var (
		logPrefix = "[artemis.ArtemisUsecase.GetStudents]"
		log       = logger.Log

		students []*model.StudentSimple
	)

	studentsDB, studentCount, err := u.artemisRepo.GetStudents(data)
	if err != nil {
		log.Printf("%v error calling artemisRepo.GetStudents. err: %v", logPrefix, err)
		return nil, -1, err
	}

	var mobilePhone string
	for _, s := range studentsDB {
		// TODO: make sure to sanitize mobile phone when inserting data to db, otherwise s[1:] operation will panic!
		if s.MobilePhoneDE.String != "" {
			mobilePhone = "+49 " + s.MobilePhoneDE.String[1:]
		} else {
			mobilePhone = "+62 " + s.MobilePhone.String[1:]
		}

		student := &model.StudentSimple{
			StudentID:    s.StudentID,
			Name:         s.GivenName + " " + s.Surname.String,
			SGUMajor:     s.SGUMajor,
			SGUEmail:     s.SGUEmail.String,
			MobilePhone:  mobilePhone,
			ExchangeYear: s.ExchangeYear.Int16,
		}

		students = append(students, student)
	}

	return students, studentCount, nil
}
