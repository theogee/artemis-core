package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/theogee/artemis-core/pkg/logger"
)

type (
	Config struct {
		Service  Service  `json:"service"`
		Database Database `json:"database"`
		Cache    Cache    `json:"cache"`
		API      API      `json:"api"`
	}

	Service struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	}

	Database struct {
		Dbname   string `json:"dbname"`
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int32  `json:"port"`
		SSLMode  string `json:"sslmode"`
	}

	Cache struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Password string `json:"password"`
		DB       int    `json:"db"`
	}

	API struct {
		AdminAuthSessionExpiration int    `json:"admin_auth_session_expiration"`
		AdminAuthCookieName        string `json:"admin_auth_cookie_name"`
	}
)

func Load(filePath string) *Config {
	var (
		log       = logger.Log
		logPrefix = "[config.Load]"

		cfg Config
	)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("%v failed to open JSON config file at: %s. err: %v", logPrefix, filePath, err)
	}

	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("%v failed to read JSON config file. err: %v", logPrefix, err)
	}

	err = json.Unmarshal(byteData, &cfg)
	if err != nil {
		log.Fatalf("%v failed to unmarshal JSON config. err: %v", logPrefix, err)
	}

	return &cfg
}
