package factory

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/gocanto/go-db-seeding/model"
)

func CreateFakeUser() model.User {
	user := model.User{}

	err := faker.FakeData(&user)

	if err != nil {
		fmt.Println(err)
	}

	return user
}
