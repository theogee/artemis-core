package artemis

import (
	"fmt"

	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func (r *ArtemisRepo) InsertStudents(students []*model.Student) error {
	var (
		logPrefix = "[artemis.ArtemisRepo.InsertStudents]"
		log       = logger.Log
	)

	if len(students) == 0 {
		log.Printf("%v aborting insert operation due to empty list", logPrefix)
		return nil
	}

	q := InsertStudentsQuery + " "

	for i, s := range students {
		if s == nil {
			log.Printf("%v nil student found", logPrefix)
			continue
		}

		hashPwd, err := bcrypt.GenerateFromPassword([]byte(s.Password), 10)
		if err != nil {
			log.Printf("%v error encrypting password. err: %v", logPrefix, err)
			continue
		}

		SGUMajorID := model.SGUMajors[s.SGUMajor]
		FHDepartmentID := model.FHDepartments[s.FHDepartment]

		q += fmt.Sprintf("('%v','%v','%v','%v','%v',%v,'%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v')", s.GivenName, s.Surname, s.Gender, SGUMajorID, FHDepartmentID, s.StudentID, s.DateOfBirth, s.CityOfBirth, s.PassportNumber, s.DateOfIssue, s.DateOfExpiry, s.IssuingOffice, s.PrivateEmail, s.SGUEmail, s.Username, string(hashPwd), s.FHEmail, s.IBAN, s.MobilePhone, s.MobilePhoneDE)

		if i != len(students)-1 {
			q += ","
		}
	}

	_, err := r.db.Conn.Exec(q)
	if err != nil {
		log.Printf("%v error inserting student data to database. err: %v", logPrefix, err)
		return err
	}

	log.Printf("%v student data has been inserted to database", logPrefix)

	return nil
}
