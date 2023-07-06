package artemis

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
)

func (r *ArtemisRepo) ParseStudentCSV(path string) ([]*model.StudentCSV, error) {
	var (
		logPrefix = "[artemis.ArtemisRepo.ParseStudentCSV]"
		log       = logger.Log

		data []*model.StudentCSV
	)

	f, err := os.Open(path)
	if err != nil {
		log.Printf("%v error opening csv file. err: %v", logPrefix, err)
		return nil, err
	}
	defer f.Close()

	err = gocsv.UnmarshalFile(f, &data)
	if err != nil {
		log.Printf("%v error unmarshalling csv file. err: %v", logPrefix, err)
		return nil, err
	}

	return data, nil
}
