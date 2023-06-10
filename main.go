package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func main() {

	db, err := gorm.Open(mysql.Open("root:12345678@/go_biz_admin"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	fmt.Println("database init...", db)

	r := gin.Default()

	// Define your handlers
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
