package artemis

import (
	"database/sql"

	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
)

func (r *ArtemisRepo) GetAdminByUsername(username string) (*model.Admin, error) {
	var (
		logPrefix = "[artemis.ArtemisRepo.GetAdminByUsername]"
		log       = logger.Log

		admin model.Admin
	)

	err := r.db.Conn.QueryRowx(GetAdminByUsernameQuery, username).StructScan(&admin)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("%v error admin with the username: %v can't be found. err: %v", logPrefix, username, err)
			return nil, nil
		}

		log.Printf("%v error fetching data from database. err: %v", logPrefix, err)
		return nil, err
	}

	return &admin, nil
}
