package artemis

import (
	"database/sql"
	"fmt"

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

func (r *ArtemisRepo) GetStudents(data *model.GetStudentsRequest) ([]*model.Student, error) {
	var (
		logPrefix = "[artemis.ArtemisRepo.GetStudents]"
		log       = logger.Log

		students []*model.Student

		args = []any{}

		offset = (data.Page - 1) * data.Limit

		count = 1
	)

	q := GetStudentsQuery

	if data.Name != "" {
		q += fmt.Sprintf(" AND CONCAT(s.given_name, ' ', s.surname) ILIKE $%v", count)
		args = append(args, "%"+data.Name+"%")
		count++
	}

	if data.ExchangeYear != 0 {
		q += fmt.Sprintf(" AND s.exchange_year = $%v", count)
		args = append(args, data.ExchangeYear)
		count++
	}

	if data.SGUMajorID != 0 {
		q += fmt.Sprintf(" AND s.sgu_major_id = $%v", count)
		args = append(args, data.SGUMajorID)
		count++
	}

	if data.StudentID != 0 {
		q += fmt.Sprintf(" AND s.student_id = $%v", count)
		args = append(args, data.StudentID)
		count++
	}

	q += fmt.Sprintf(" LIMIT $%v OFFSET $%v", count, count+1)
	args = append(args, data.Limit, offset)

	err := r.db.Conn.Select(&students, q, args...)
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
