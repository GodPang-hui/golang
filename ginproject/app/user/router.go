package user

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/topgoer", helloHandler)
	e.GET("/ping", pangHandler)
	e.GET("/user/:name/*action", userNameActionHandler)
	e.GET("/user", userHandler)
}
