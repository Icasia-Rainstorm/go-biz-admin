package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mousepotato/go-biz-admin/database"
	"github.com/mousepotato/go-biz-admin/models"
	"net/http"
	"strconv"
)

func AllProducts(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	ret := models.Paginate(database.DB, &models.Product{}, page)
	c.JSON(http.StatusOK, ret)
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
