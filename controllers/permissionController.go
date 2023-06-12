package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mousepotato/go-biz-admin/database"
	"github.com/mousepotato/go-biz-admin/models"
	"net/http"
)

func AllPermissions(c *gin.Context) {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	c.JSON(http.StatusOK, permissions)
}
