package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mousepotato/go-biz-admin/controllers"
	"github.com/mousepotato/go-biz-admin/middlewares"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// public handlers
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)

	r.Use(middlewares.IsAuthenticated)

	// User Handlers
	r.GET("/api/user", controllers.User)
	r.POST("/api/logout", controllers.Logout)
	r.PUT("/api/users/info", controllers.UpdateInfo)
	r.PUT("/api/users/password", controllers.UpdatePassword)

	r.GET("/api/users", controllers.AllUsers)
	r.POST("/api/users", controllers.CreateUser)
	r.GET("/api/users/:id", controllers.GetUser)
	r.PUT("/api/users/:id", controllers.UpdateUser)
	r.DELETE("/api/users/:id", controllers.DeleteUser)

	r.GET("/api/roles", controllers.AllRoles)
	r.POST("/api/roles", controllers.CreateRole)
	r.GET("/api/roles/:id", controllers.GetRole)
	r.PUT("/api/roles/:id", controllers.UpdateRole)
	r.DELETE("/api/roles/:id", controllers.DeleteRole)

	r.GET("/api/permissions", controllers.AllPermissions)

	r.GET("/api/products", controllers.AllProducts)
	r.POST("/api/products", controllers.CreateProduct)
	r.GET("/api/products/:id", controllers.GetProduct)
	r.PUT("/api/products/:id", controllers.UpdateProduct)
	r.DELETE("/api/products/:id", controllers.DeleteProduct)

	r.POST("/api/upload", controllers.Upload)
	r.Static("/api/uploads", "./uploads")

	return r
}
