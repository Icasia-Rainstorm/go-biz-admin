package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	form, err := c.MultipartForm()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "file upload error"})
		return
	}

	files := form.File["image"]
	filename := ""

	for _, file := range files {
		filename = file.Filename
		if err := c.SaveUploadedFile(file, "./uploads/"+filename); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "file upload error"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"url": "http://localhost:8080/api/uploads/" + filename,
	})
}
