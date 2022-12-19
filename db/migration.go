package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate() {
	db, err := GetDbConnection()

	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		panic(err)
	}

	// NOTE: relative path to folder
	m, err := migrate.NewWithDatabaseInstance(
		"file://./schema",
		"postgres", driver)

	if err != nil {
		panic(err)
	}

	err = m.Up()

	if err != nil {
		fmt.Println(err)
	}
}
