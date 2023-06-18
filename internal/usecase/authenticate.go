package artemis

import (
	"github.com/theogee/artemis-core/pkg/logger"
)

func (u *ArtemisUsecase) Authenticate(sid string) (string, error) {
	var (
		logPrefix = "[artemis.ArtemisUsecase.Authenticate]"
		log       = logger.Log
	)

	uid, err := u.artemisRepo.GetCache(sid)
	if err != nil {
		log.Printf("%v error calling artemisRepo.GetCache. err: %v", logPrefix, err)
		return "", err
	}

	return uid, nil
}
