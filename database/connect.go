package database

import (
	"fmt"
	"github.com/mousepotato/go-biz-admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:12345678@/go_biz_admin"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	fmt.Println("database init...", database)
	DB = database
	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})
}
