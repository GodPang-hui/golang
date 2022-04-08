package routers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

func MiddleHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		ip := c.Request.RemoteAddr
		fmt.Println("ip:host = " + ip)
		c.Next()
		// 统计时间
		since := time.Since(start)
		fmt.Println("程序用时：", since)
	}
}

// 初始化
func Init() *gin.Engine {
	r := gin.Default()
	r.Use(MiddleHandler())
	r.MaxMultipartMemory = 8 << 20
	fmt.Println(r.MaxMultipartMemory)
	for _, opt := range options {
		opt(r)
	}
	return r
}
