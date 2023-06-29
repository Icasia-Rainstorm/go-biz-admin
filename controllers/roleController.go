package controllers

import (
	"encoding/json"
	"fmt"
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
	var roleDto map[string]interface{}

	val, _ := c.GetRawData()
	fmt.Println(string(val))
	err := json.Unmarshal(val, &roleDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "input role data unmarshal error"},
		)
	}

	list := roleDto["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))

	for idx, permissionId := range list {
		id, _ := permissionId.(float64)
		permissions[idx] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:        roleDto["name"].(string),
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
	var roleDto map[string]interface{}

	val, _ := c.GetRawData()
	fmt.Println(string(val))
	err := json.Unmarshal(val, &roleDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "input role data unmarshal error"},
		)
	}
	list := roleDto["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))

	for idx, permissionId := range list {
		id, _ := permissionId.(float64)
		permissions[idx] = models.Permission{
			Id: uint(id),
		}
	}

	var result struct{} // var result interface{} this will throw error as no zero value!
	database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)

	role := models.Role{
		Id:          uint(id),
		Name:        roleDto["name"].(string),
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

	var result struct{}
	database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)

	database.DB.Delete(&role)
	c.JSON(http.StatusOK, gin.H{"message": "role delete successfully"})
}
