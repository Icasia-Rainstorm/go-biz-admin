package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/mousepotato/go-biz-admin/util"
	"net/http"
)

func IsAuthenticated(c *gin.Context) {
	cookie, _ := c.Cookie("jwt")
	if _, err := util.ParseJwt(cookie); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthenticated"},
		)
		return
	}
	c.Next()
}
