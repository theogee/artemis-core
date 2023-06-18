package artemis

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func (u *ArtemisUsecase) LoginAsAdmin(data *model.LoginAsAdminRequest) (string, error) {
	var (
		logPrefix = "[artemis.ArtemisUsecase.LoginAsAdmin]"
		log       = logger.Log
	)

	admin, err := u.artemisRepo.GetAdminByUsername(data.Username)
	if err != nil {
		log.Printf("%v error calling artemisRepo.GetAdminByUsername. err: %v", logPrefix, err)
		return "", err
	}

	if admin == nil {
		log.Printf("%v admin with username: %v can't be found", logPrefix, data.Username)
		return "", errors.New(model.IncorrectCredential)
	}

	// match password
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(data.Password))
	if err != nil {
		log.Printf("%v incorrect password. err: %v", logPrefix, err)
		return "", errors.New(model.IncorrectCredential)
	}

	// TODO: generate sid and set sid:adminID to redis
	sid := uuid.New().String()

	err = u.artemisRepo.SetCache(sid, admin.AdminID, time.Second*time.Duration(u.cfg.API.AdminAuthSessionExpiration))
	if err != nil {
		log.Printf("%v error calling artemisRepo.SetCache. err: %v", logPrefix, err)
		return "", err
	}

	return sid, nil
}
