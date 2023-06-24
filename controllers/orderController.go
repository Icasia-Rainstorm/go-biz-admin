package controllers

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/mousepotato/go-biz-admin/database"
	"github.com/mousepotato/go-biz-admin/models"
	"net/http"
	"os"
	"strconv"
)

func AllOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	ret := models.Paginate(database.DB, &models.Order{}, page)
	c.JSON(http.StatusOK, ret)
}

func Export(c *gin.Context) {
	filePath := "./csv/orders.csv"

	if err := CreateFile(filePath); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "export csv file error"})
		return
	}

	c.File(filePath)
}

func CreateFile(filePath string) error {
	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var orders []models.Order

	database.DB.Preload("OrderItems").Find(&orders)

	writer.Write([]string{
		"ID", "Name", "Email", "Product Title", "Price", "Quantity",
	})

	for _, order := range orders {
		data := []string{
			strconv.Itoa(int(order.Id)),
			order.FirstName + " " + order.LastName,
			order.Email,
			"",
			"",
			"",
		}

		if err := writer.Write(data); err != nil {
			return err
		}

		for _, orderItem := range order.OrderItems {
			data := []string{
				"",
				"",
				"",
				orderItem.ProductTitle,
				strconv.Itoa(int(orderItem.Price)),
				strconv.Itoa(int(orderItem.Quantity)),
			}

			if err := writer.Write(data); err != nil {
				return err
			}
		}
	}

	return nil
}

type Sales struct {
	Date string `json:"date"`
	Sum  string `json:"sum"`
}

func Chart(c *gin.Context) {
	var sales []Sales

	database.DB.Raw(`
		SELECT DATE_FORMAT(o.created_at, '%Y-%m-%d') as date, SUM(oi.price * oi.quantity) as sum
		FROM orders o
		JOIN order_items oi on o.id = oi.order_id
		GROUP BY date
	`).Scan(&sales)

	c.JSON(http.StatusOK, sales)
}
