package artemis

import (
	"database/sql"

	"github.com/theogee/artemis-core/pkg/logger"
)

func (r *ArtemisRepo) GetExchangeYear() ([]int, error) {
	var (
		logPrefix = "[artemis.ArtemisRepo.GetExchangeYear]"
		log       = logger.Log

		exchangeYear []int
	)

	err := r.db.Conn.Select(&exchangeYear, GetExchangeYearQuery)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("%v error there is no data in students table. err: %v", logPrefix, err)
			return exchangeYear, nil
		}

		log.Printf("%v error fetching data from database. err: %v", logPrefix, err)
		return nil, err
	}

	return exchangeYear, nil
}
