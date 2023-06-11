package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mousepotato/go-biz-admin/controllers"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define your handlers
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// User Handlers
	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)

	return r
}
