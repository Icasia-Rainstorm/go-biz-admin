package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mousepotato/go-biz-admin/database"
	"github.com/mousepotato/go-biz-admin/models"
	"math"
	"net/http"
	"strconv"
)

func AllProducts(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64

	var products []models.Product

	database.DB.Offset(offset).Limit(limit).Find(&products)
	database.DB.Model(&models.Product{}).Count(&total)

	c.JSON(http.StatusOK, gin.H{
		"data": products,
		"meta": gin.H{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "invalid user JSON file"},
		)
		return
	}

	database.DB.Create(&product)
	c.JSON(http.StatusOK, product)
}

func GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Find(&product)

	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	product := models.Product{
		Id: uint(id),
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "invalid user JSON file"},
		)
		return
	}

	database.DB.Model(&product).Updates(product)

	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "user delete successfully"})
}
