package main

import (
	"github.com/bxcodec/faker/v3"
	"github.com/mousepotato/go-biz-admin/database"
	"github.com/mousepotato/go-biz-admin/models"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		user := models.User{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Email:     faker.Email(),
			RoleId:    1,
		}

		user.SetPassword("1234")

		database.DB.Create(&user)
	}
}
