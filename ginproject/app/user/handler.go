package user

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com!",
	})
}

func pangHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func userNameActionHandler(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	//截取/
	action = strings.Trim(action, "/")
	// c.String(http.StatusOK, name+" is "+action)

	c.JSON(http.StatusOK, gin.H{
		"message": name + " is " + action,
	})
}

func userHandler(c *gin.Context) {
	//指定默认值
	//http://localhost:8080/user 才会打印出来默认的值
	name := c.DefaultQuery("name", "枯藤")
	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
}
