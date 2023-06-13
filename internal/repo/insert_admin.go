package artemis

import (
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
)

func (r *ArtemisRepo) InsertAdmin(data *model.Admin) error {
	var (
		logPrefix = "[artemis.ArtemisRepo.InsertAdmin]"
		log       = logger.Log
	)

	_, err := r.db.Conn.Exec(InsertAdminQuery, data.Username, data.Password, data.Email)
	if err != nil {
		log.Printf("%v error inserting admin data to database. err: %v", logPrefix, err)
		return err
	}

	log.Printf("%v new admin with username: %v has been inserted to database", logPrefix, data.Username)

	return nil
}
