package db

import "database/sql"

func GetDbConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:pwd@localhost:5432/gus-db?sslmode=disable")

	if err != nil {
		return nil, err
	}

	return db, nil
}
