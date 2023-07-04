package artemis

import (
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
)

func (u *ArtemisUsecase) GetSGUMajors() ([]*model.SGUMajor, error) {
	var (
		logPrefix = "[artemis.ArtemisRepo.GetSGUMajors]"
		log       = logger.Log
	)

	SGUMajors, err := u.artemisRepo.GetSGUMajors()
	if err != nil {
		log.Printf("%v error calling artemisRepo.GetSGUMajors. err: %v", logPrefix, err)
		return nil, err
	}

	return SGUMajors, nil
}
