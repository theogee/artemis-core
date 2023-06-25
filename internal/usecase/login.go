package artemis

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func (u *ArtemisUsecase) Login(data *model.LoginRequest) (string, error) {
	var (
		logPrefix = "[artemis.ArtemisUsecase.Login]"
		log       = logger.Log

		password string
		id       uint32
	)

	admin, err := u.artemisRepo.GetAdminByUsername(data.Username)
	if err != nil {
		log.Printf("%v error calling artemisRepo.GetAdminByUsername. err: %v", logPrefix, err)
		return "", err
	}

	if admin == nil {
		log.Printf("%v admin with username: %v can't be found", logPrefix, data.Username)

		student, err := u.artemisRepo.GetStudentByUsername((data.Username))
		if err != nil {
			log.Printf("%v error calling artemisRepo.GetStudentByUsername. err: %v", logPrefix, err)
			return "", err
		}

		if student == nil {
			log.Printf("%v student with username: %v can't be found", logPrefix, data.Username)
			return "", errors.New(model.IncorrectCredential)
		}

		password = student.Password
		id = student.StudentID
	} else {
		password = admin.Password
		id = admin.AdminID
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(data.Password))
	if err != nil {
		log.Printf("%v incorrect password. err: %v", logPrefix, err)
		return "", errors.New(model.IncorrectCredential)
	}

	sid := uuid.New().String()

	err = u.artemisRepo.SetCache(sid, id, time.Second*time.Duration(u.cfg.API.AuthSessionExpiration))
	if err != nil {
		log.Printf("%v error calling artemisRepo.SetCache. err: %v", logPrefix, err)
		return "", err
	}

	return sid, nil
}
