package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mousepotato/go-biz-admin/database"
	"github.com/mousepotato/go-biz-admin/models"
	"net/http"
)

func Register(c *gin.Context) {
	var data map[string]string

	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Input data is not JSON format"})
		return
	}

	if data["password"] != data["password_confirm"] {
		fmt.Println("password does not match...")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "password does not match"})
		return
	}
	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var data map[string]string

	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Input data is not JSON format"})
		return
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "incorrect password"})
		return
	}

	c.JSON(http.StatusOK, user)
}
