package db

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	dialect = "pgx"
)

func RunMigrate(command, dbString, dir string, args ...string) error {
	db, err := goose.OpenDBWithDriver(dialect, dbString)
	if err != nil {
		log.Fatalf(err.Error())
		return err
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf(err.Error())
		}
	}()

	if err := goose.Run(command, db, dir, args...); err != nil {
		log.Fatalf("migrate %v: %v", command, err)
		return err
	}

	return nil
}
