package artemis

import (
	"database/sql"

	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
)

func (r *ArtemisRepo) GetSGUMajors() ([]*model.SGUMajor, error) {
	var (
		logPrefix = "[artemis.ArtemisRepo.GetSGUMajors]"
		log       = logger.Log

		SGUMajors []*model.SGUMajor
	)

	err := r.db.Conn.Select(&SGUMajors, GetSGUMajorsQuery)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("%v error there is no data in sgu_majors table. err: %v", logPrefix, err)
			return SGUMajors, nil
		}

		log.Printf("%v error fetching data from database. err: %v", logPrefix, err)
		return nil, err
	}

	return SGUMajors, nil
}
