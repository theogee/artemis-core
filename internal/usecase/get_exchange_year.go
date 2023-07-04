package artemis

import "github.com/theogee/artemis-core/pkg/logger"

func (u *ArtemisUsecase) GetExchangeYear() ([]int, error) {
	var (
		logPrefix = "[artemis.ArtemisUsecase.GetExchangeYear]"
		log       = logger.Log
	)

	exchangeYear, err := u.artemisRepo.GetExchangeYear()
	if err != nil {
		log.Printf("%v error calling artemisRepo.GetExchangeYear. err: %v", logPrefix, err)
		return nil, err
	}

	return exchangeYear, nil
}
