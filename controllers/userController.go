package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mousepotato/go-biz-admin/database"
	"github.com/mousepotato/go-biz-admin/models"
	"net/http"
	"strconv"
)

func AllUsers(c *gin.Context) {
	var users []models.User

	database.DB.Find(&users)

	c.JSON(http.StatusOK, users)

}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "invalid user JSON file"},
		)
		return
	}

	user.SetPassword("1234")

	database.DB.Create(&user)

	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Find(&user)

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "invalid user JSON file"},
		)
		return
	}

	database.DB.Model(&user).Updates(user)

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "user delete successfully"})
}
