package artemis

import (
	"encoding/json"
	"errors"
	"strconv"
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
		userType string
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
		userType = "student"
	} else {
		password = admin.Password
		id = admin.AdminID
		userType = "admin"
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(data.Password))
	if err != nil {
		log.Printf("%v incorrect password. err: %v", logPrefix, err)
		return "", errors.New(model.IncorrectCredential)
	}

	sid := uuid.New().String()

	c := &model.UserCache{
		UID:      strconv.FormatUint(uint64(id), 10),
		UserType: userType,
	}

	cstr, err := json.Marshal(c)
	if err != nil {
		log.Printf("%v error marshaling cache data. err: %v", logPrefix, err)
		return "", err
	}

	err = u.artemisRepo.SetCache(sid, string(cstr), time.Second*time.Duration(u.cfg.API.AuthSessionExpiration))
	if err != nil {
		log.Printf("%v error calling artemisRepo.SetCache. err: %v", logPrefix, err)
		return "", err
	}

	return sid, nil
}
