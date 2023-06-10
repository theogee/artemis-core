package main

import (
	"github.com/theogee/artemis-core/pkg/logger"
)

func main() {
	logger.Setup("main.log")
	
	var (
		logPrefix = "[main]"
		log = logger.Log
	)
	
	log.Printf("%v hello from artemis-core!", logPrefix)

	logger.Close()
}