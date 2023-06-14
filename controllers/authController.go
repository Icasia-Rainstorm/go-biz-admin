package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mousepotato/go-biz-admin/database"
	"github.com/mousepotato/go-biz-admin/models"
	"github.com/mousepotato/go-biz-admin/util"
	"net/http"
	"strconv"
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
		RoleId:    1, // set as Admin
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

	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.SetCookie("jwt", token, 3600, "", "", false, true)

	// c.JSON(http.StatusOK, token)
	c.JSON(http.StatusOK, user)
}

func User(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		fmt.Println("cookie not set")
	}

	id, _ := util.ParseJwt(cookie)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	c.JSON(http.StatusOK, user)
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", 0, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}

func UpdateInfo(c *gin.Context) {
	var data map[string]string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Input data is not JSON format"})
		return
	}
	cookie, err := c.Cookie("jwt")
	if err != nil {
		fmt.Println("cookie not set")
	}
	id, _ := util.ParseJwt(cookie)
	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id:        uint(userId),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}
	database.DB.Model(&user).Updates(user)

	c.JSON(http.StatusOK, user)
}

func UpdatePassword(c *gin.Context) {
	var data map[string]string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Input data is not JSON format"})
		return
	}

	if data["password"] != data["password_confirm"] {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "passwords do not match"})
	}

	cookie, err := c.Cookie("jwt")
	if err != nil {
		fmt.Println("cookie not set")
	}

	id, _ := util.ParseJwt(cookie)
	userId, _ := strconv.Atoi(id)
	user := models.User{
		Id: uint(userId),
	}
	user.SetPassword(data["password"])
	database.DB.Model(&user).Updates(user)

	c.JSON(http.StatusOK, user)
}
