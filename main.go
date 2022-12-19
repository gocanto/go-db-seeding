package main

import (
	"fmt"
	localDB "github.com/gocanto/go-db-seeding/db"
	"github.com/gocanto/go-db-seeding/db/factory"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	db, err := localDB.GetDbConnection()

	if err != nil {
		panic(err)
	}

	user := factory.CreateFakeUser()

	lastInsertId := 0

	err = db.QueryRow(
		"INSERT INTO users (username, handle) VALUES ($1, $2) RETURNING id;",
		user.Username,
		user.Handle,
	).Scan(&lastInsertId)

	if err != nil {
		panic(err)
	}

	fmt.Println(lastInsertId)
}
