package main

import (
	"github.com/bxcodec/faker/v3"
	"github.com/mousepotato/go-biz-admin/database"
	"github.com/mousepotato/go-biz-admin/models"
	"math/rand"
	"time"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		var orderItems []models.OrderItem

		for j := 0; j < rand.Intn(5); j++ {
			price := float32(rand.Intn(90) + 10)
			qty := uint(rand.Intn(5))

			orderItems = append(orderItems, models.OrderItem{
				OrderId:      uint(rand.Intn(5)),
				ProductTitle: faker.Word(),
				Price:        price,
				Quantity:     qty,
			})
		}

		database.DB.Create(&models.Order{
			FirstName:  faker.FirstName(),
			LastName:   faker.LastName(),
			Email:      faker.Email(),
			OrderItems: orderItems,
			UpdatedAt:  time.Unix(faker.RandomUnixTime(), 0).Format("2006-01-02 15:04:05"),
			CreatedAt:  time.Unix(faker.RandomUnixTime(), 0).Format("2006-01-02 15:04:05"),
			Total:      float32(rand.Intn(90) + 10),
		})
	}
}
