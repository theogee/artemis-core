package artemis

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/theogee/artemis-core/pkg/logger"
)

func (r *ArtemisRepo) SaveFile(file multipart.File, header *multipart.FileHeader, destination string) error {
	var (
		logPrefix = "[artemis.ArtemisRepo.SaveFile]"
		log       = logger.Log
	)

	f, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("%v error creating destination file. err: %v", logPrefix, err)
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		log.Printf("%v error copying temp file to destination file. err: %v", logPrefix, err)
		return err
	}

	log.Printf("%v file has been saved. path: %v", logPrefix, destination)
	return nil
}
