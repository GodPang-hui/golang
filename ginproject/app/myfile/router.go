package myfile

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.POST("/form", formHandler)
	e.POST("/upload", uploadHandler)
	e.POST("/uploads", uploadsHandler)
}
