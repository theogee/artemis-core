package artemis

import (
	"github.com/theogee/artemis-core/pkg/logger"
)

func (u *ArtemisUsecase) Logout(sid string) error {
	var (
		logPrefix = "[artemis.ArtemisUsecase.Logout]"
		log       = logger.Log
	)

	err := u.artemisRepo.DeleteCache(sid)
	if err != nil {
		log.Printf("%v error calling artemisRepo.DeleteCache. err: %v", logPrefix, err)
		return err
	}

	return nil
}
