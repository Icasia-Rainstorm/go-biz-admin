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

// RoleCreateDTO : make sure field name in Upper case for JSON bind
type RoleCreateDTO struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

func CreateRole(c *gin.Context) {
	var roleDto RoleCreateDTO

	if err := c.ShouldBindJSON(&roleDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "invalid role JSON file"},
		)
		return
	}

	permissions := make([]models.Permission, len(roleDto.Permissions))

	for idx, permissionId := range roleDto.Permissions {
		id, _ := strconv.Atoi(permissionId)
		permissions[idx] = models.Permission{Id: uint(id)}
	}

	role := models.Role{
		Name:        roleDto.Name,
		Permissions: permissions,
	}
	database.DB.Create(&role)

	c.JSON(http.StatusOK, role)
}

func GetRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Preload("Permissions").Find(&role)

	c.JSON(http.StatusOK, role)
}

func UpdateRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var roleDto RoleCreateDTO

	if err := c.ShouldBindJSON(&roleDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "invalid role JSON file"},
		)
		return
	}

	permissions := make([]models.Permission, len(roleDto.Permissions))

	for idx, permissionId := range roleDto.Permissions {
		id, _ := strconv.Atoi(permissionId)
		permissions[idx] = models.Permission{Id: uint(id)}
	}
	var result struct{} // var result interface{} this will throw error as no zero value!
	database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)

	role := models.Role{
		Id:          uint(id),
		Name:        roleDto.Name,
		Permissions: permissions,
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
