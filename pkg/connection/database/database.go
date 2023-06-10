package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/theogee/artemis-core/pkg/config"
	"github.com/theogee/artemis-core/pkg/logger"
)

type (
	Database struct {
		Conn   *sqlx.DB
		Config *config.Config
	}
)

func NewDatabase(cfg *config.Config) *Database {
	var (
		logPrefix = "[database.NewDatabase]"
		log       = logger.Log
	)

	d := &Database{
		Config: cfg,
	}

	err := d.connect()
	if err != nil {
		log.Fatalf("%v error connecting to artemis-core-postgres. err: %v", logPrefix, err)
	}

	return d
}

func (d *Database) connect() error {
	var (
		logPrefix = "[database.connect]"
		log       = logger.Log

		connStr = fmt.Sprintf("dbname=%v user=%v password=%v host=%v port=%v sslmode=%v", d.Config.Database.Dbname, d.Config.Database.User, d.Config.Database.Password, d.Config.Database.Host, d.Config.Database.Port, d.Config.Database.SSLMode)
	)

	conn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Printf("%v error connecting to artemis-core-postgres. err: %v", logPrefix, err)
		return err
	}

	d.Conn = conn
	log.Printf("%v connected to artemis-core-postgres", logPrefix)

	return nil
}
