package parser

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
)

func ParseCSV(path string) []*model.Student {
	var (
		logPrefix = "[parser.ParseCSV]"
		log       = logger.Log
	)

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("%v error opening csv file. err: %v", logPrefix, err)
	}
	defer f.Close()

	data := []*model.Student{}

	err = gocsv.UnmarshalFile(f, &data)
	if err != nil {
		log.Fatalf("%v error unmarshaling csv file. err: %v", logPrefix, err)
	}

	return data
}
