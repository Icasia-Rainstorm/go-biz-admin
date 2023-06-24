package middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mousepotato/go-biz-admin/database"
	"github.com/mousepotato/go-biz-admin/models"
	"github.com/mousepotato/go-biz-admin/util"
	"strconv"
)

func IsAuthorized(c *gin.Context, page string) error {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		fmt.Println("cookie not set")
	}

	Id, err := util.ParseJwt(cookie)
	fmt.Println(Id)
	if err != nil {
		return err
	}
	userId, _ := strconv.Atoi(Id)

	user := models.User{
		Id: uint(userId),
	}

	database.DB.Preload("Role").Find(&user)

	role := models.Role{
		Id: user.RoleId,
	}

	database.DB.Preload("Permissions").Find(&role)

	fmt.Println(role.Permissions)
	if c.Request.Method == "GET" {
		for _, permission := range role.Permissions {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permissions {
			fmt.Println(permission.Name)
			if permission.Name == "edit_"+page {
				return nil
			}
		}
	}

	return errors.New("unauthorized")
}
