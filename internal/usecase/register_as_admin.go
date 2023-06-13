package artemis

import (
	"errors"

	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func (u *ArtemisUsecase) RegisterAsAdmin(data *model.RegisterAsAdminRequest) error {
	var (
		logPrefix = "[artemis.ArtemisUsecase.RegisterAsAdmin]"
		log       = logger.Log
	)

	if data == nil {
		log.Printf("%v error data parameter is nil", logPrefix)
		return errors.New("data parameter is nil")
	}

	admin, err := u.artemisRepo.GetAdminByUsername(data.Username)
	if err != nil {
		log.Printf("%v error calling artemisRepo.GetAdminByUsername. err: %v", logPrefix, err)
		return err
	}

	// admin with the same username exist
	if admin != nil {
		log.Printf("%v error can't register admin. err: username already exist", logPrefix)
		return errors.New(model.UsernameAlreadyExist)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		log.Printf("%v error calling bcrypt.GenerateFromPassword. err: %v", logPrefix, err)
		return err
	}

	newAdmin := &model.Admin{
		Username: data.Username,
		Password: string(hashPassword),
		Email:    data.Email,
	}

	err = u.artemisRepo.InsertAdmin(newAdmin)
	if err != nil {
		log.Printf("%v error calling artemisRepo.InsertAdmin. err: %v", logPrefix, err)
		return err
	}

	return nil
}
