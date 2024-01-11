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
	var sguEmail string
	for _, s := range studentsDB {
		// TODO: make sure to sanitize mobile phone when inserting data to db, otherwise s[1:] operation will panic!
		// if s.MobilePhoneDE.String != "" {
		// 	mobilePhone = "+49 " + s.MobilePhoneDE.String[1:]
		// } else {
		// 	mobilePhone = "+62 " + s.MobilePhone.String[1:]
		// }

		if s.MobilePhoneDE.String != "" {
			mobilePhone = s.MobilePhoneDE.String
		} else if s.MobilePhone.String != "" {
			mobilePhone = s.MobilePhone.String
		} else {
			mobilePhone = "-"
		}

		if s.SGUEmail.String != "" {
			sguEmail = s.SGUEmail.String
		} else {
			sguEmail = "-"
		}

		student := &model.StudentSimple{
			StudentID:    s.StudentID,
			Name:         s.GivenName + " " + s.Surname.String,
			SGUMajor:     s.SGUMajor,
			SGUEmail:     sguEmail,
			MobilePhone:  mobilePhone,
			ExchangeYear: s.ExchangeYear.Int16,
		}

		students = append(students, student)
	}

	return students, studentCount, nil
}

func (u *ArtemisUsecase) GetStudentByID(studentID int64) (*model.StudentDetail, error) {
	var (
		logPrefix = "[artemis.ArtemisUsecase.GetStudentByID]"
		log       = logger.Log

		studentDB *model.Student
		student   *model.StudentDetail
	)

	studentDB, err := u.artemisRepo.GetStudentByID(studentID)
	if err != nil {
		log.Printf("%v error calling artemisRepo.GetStudentByID. err: %v", logPrefix, err)
		return nil, err
	}

	if studentDB == nil {
		return nil, nil
	}

	student = &model.StudentDetail{
		GivenName:                 studentDB.GivenName,
		Surname:                   studentDB.Surname.String,
		Username:                  studentDB.Username,
		Gender:                    studentDB.Gender.String,
		SGUMajor:                  studentDB.SGUMajor,
		SGUMajorInitial:           studentDB.SGUMajorCode,
		FHDepartment:              studentDB.FHDepartment,
		StudentID:                 studentDB.StudentID,
		DateOfBirth:               studentDB.DateOfBirth.String,
		CityOfBirth:               studentDB.CityOfBirth.String,
		PassportNumber:            studentDB.PassportNumber.String,
		DateOfIssue:               studentDB.DateOfIssue.String,
		DateOfExpiry:              studentDB.DateOfExpiry.String,
		IssuingOffice:             studentDB.IssuingOffice.String,
		PrivateEmail:              studentDB.PrivateEmail.String,
		SGUEmail:                  studentDB.SGUEmail.String,
		FHEmail:                   studentDB.FHEmail.String,
		IBAN:                      studentDB.IBAN.String,
		MobilePhone:               studentDB.MobilePhone.String,
		MobilePhoneDE:             studentDB.MobilePhoneDE.String,
		CurrentAddress:            studentDB.CurrentAddress.String,
		CurrentPostcode:           studentDB.CurrentPostcode.String,
		CurrentCity:               studentDB.CurrentCity.String,
		CoName:                    studentDB.CoName.String,
		InternshipCompany:         studentDB.InternshipCompany.String,
		InternshipStartDate:       studentDB.InternshipStartDate.String,
		InternshipEndDate:         studentDB.InternshipEndDate.String,
		InternshipCompanyAddress:  studentDB.InternshipCompanyAddress.String,
		InternshipCompanyPostcode: studentDB.InternshipCompanyPostcode.String,
		InternshipCompanyCity:     studentDB.InternshipCompanyCity.String,
		InternshipSupervisorName:  studentDB.InternshipSupervisorName.String,
		InternshipSupervisorEmail: studentDB.InternshipSupervisorEmail.String,
		InternshipSupervisorPhone: studentDB.InternshipSupervisorPhone.String,
		ExchangeYear:              studentDB.ExchangeYear.Int16,
	}

	return student, nil
}
