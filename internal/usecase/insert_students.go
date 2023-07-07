package artemis

import (
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
)

func (u *ArtemisUsecase) InsertStudents(students []*model.StudentCSV, exchangeYear int64) error {
	var (
		logPrefix = "[artemis.ArtemisUsecase.InsertStudents]"
		log       = logger.Log
	)

	err := u.artemisRepo.InsertStudents(students, exchangeYear)
	if err != nil {
		log.Printf("%v error calling artemisRepo.InsertStudents. err: %v", logPrefix, err)
		return err
	}

	return nil
}
