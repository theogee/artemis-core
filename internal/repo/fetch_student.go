package artemis

import (
	"database/sql"

	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
)

func (r *ArtemisRepo) GetStudentByUsername(username string) (*model.Student, error) {
	var (
		logPrefix = "[artemis.ArtemisRepo.GetStudentByUsername]"
		log       = logger.Log

		student model.Student
	)

	err := r.db.Conn.QueryRowx(GetStudentByUsernameQuery, username).StructScan(&student)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("%v error student with the username: %v can't be found. err: %v", logPrefix, username, err)
			return nil, nil
		}

		log.Printf("%v error fetching data from database. err: %v", logPrefix, err)
		return nil, err
	}

	return &student, nil
}

func (r *ArtemisRepo) GetStudents(limit, offset int64) ([]*model.Student, error) {
	var (
		logPrefix = "[artemis.ArtemisRepo.GetStudents]"
		log       = logger.Log

		students []*model.Student
	)

	err := r.db.Conn.Select(&students, GetStudentsQuery, limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("%v error there is no data in students table. err: %v", logPrefix, err)
			return students, nil
		}

		log.Printf("%v error fetching data from database. err: %v", logPrefix, err)
		return nil, err
	}

	return students, nil
}
