package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mousepotato/go-biz-admin/models"
	"net/http"
)

func Register(c *gin.Context) {
	user := models.User{
		FirstName: "John",
	}
	user.LastName = "Doe"
	user.Email = "test@gmail.com"

	c.JSON(http.StatusOK, user)
}
