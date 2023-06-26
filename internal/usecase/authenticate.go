package artemis

import (
	"encoding/json"

	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
)

func (u *ArtemisUsecase) Authenticate(sid string) (string, error) {
	var (
		logPrefix = "[artemis.ArtemisUsecase.Authenticate]"
		log       = logger.Log

		c = &model.UserCache{}
	)

	cstr, err := u.artemisRepo.GetCache(sid)
	if err != nil {
		log.Printf("%v error calling artemisRepo.GetCache. err: %v", logPrefix, err)
		return "", err
	}

	err = json.Unmarshal([]byte(cstr), c)
	if err != nil {
		log.Printf("%v error unmarshalling cache into struct. err: %v", logPrefix, err)
		return "", err
	}

	return c.UID, nil
}
