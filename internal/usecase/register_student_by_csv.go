package artemis

import (
	"mime/multipart"

	"github.com/theogee/artemis-core/pkg/logger"
)

func (u *ArtemisUsecase) RegisterStudentByCSV(file multipart.File, header *multipart.FileHeader, exchangeYear int64) error {
	var (
		logPrefix = "[artemis.ArtemisUsecase.RegisterStudentByCSV]"
		log       = logger.Log

		destination = "resources/uploads/" + header.Filename
	)

	err := u.artemisRepo.SaveFile(file, header, destination)
	if err != nil {
		log.Printf("%v error calling artemisRepo.SaveFile. err: %v", logPrefix, err)
		return err
	}

	// TODO: parse the file and insert to db
	data, err := u.artemisRepo.ParseStudentCSV(destination)
	if err != nil {
		log.Printf("%v error calling artemisRepo.ParseStudentCSV. err: %v", logPrefix, err)
		return err
	}

	err = u.artemisRepo.InsertStudents(data, exchangeYear)
	if err != nil {
		log.Printf("%v error calling artemisRepo.InsertStudents. err: %v", logPrefix, err)
		return err
	}

	return nil
}
