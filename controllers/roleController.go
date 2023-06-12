package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mousepotato/go-biz-admin/database"
	"github.com/mousepotato/go-biz-admin/models"
	"net/http"
	"strconv"
)

func AllRoles(c *gin.Context) {
	var role []models.Role

	database.DB.Find(&role)

	c.JSON(http.StatusOK, role)

}

func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "invalid role JSON file"},
		)
		return
	}

	database.DB.Create(&role)

	c.JSON(http.StatusOK, role)
}

func GetRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Find(&role)

	c.JSON(http.StatusOK, role)
}

func UpdateRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	role := models.Role{
		Id: uint(id),
	}

	if err := c.ShouldBindJSON(&role); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "invalid role JSON file"},
		)
		return
	}

	database.DB.Model(&role).Updates(role)

	c.JSON(http.StatusOK, role)
}

func DeleteRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)
	c.JSON(http.StatusOK, gin.H{"message": "role delete successfully"})
}
