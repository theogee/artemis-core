package artemis

import (
	"errors"

	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
)

func (u *ArtemisUsecase) UpdateStudentByID(data *model.UpdateStudentByIDRequest) error {
	var (
		logPrefix = "[artemis.ArtemisUsecase.UpdateStudentByID]"
		log       = logger.Log
	)

	if data == nil {
		log.Printf("%v error data parameter is nil", logPrefix)
		return errors.New("data parameter is nil")
	}

	// check if studentID exist
	student, err := u.artemisRepo.GetStudentByID(data.StudentID)
	if err != nil {
		log.Printf("%v error calling artemisRepo.GetStudentByID. err: %v", logPrefix, err)
		return err
	}

	if student == nil {
		log.Printf("%v error studentID: %v doesn't exist", logPrefix, data.StudentID)
		return errors.New(model.StudentIDDoesNotExist)
	}

	err = u.artemisRepo.UpdateStudentByID(data)
	if err != nil {
		log.Printf("%v error calling artemisRepo.UpdateStudentByID. err: %v", logPrefix, err)
		return err
	}

	return nil
}
