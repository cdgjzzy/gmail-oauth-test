package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// 注册handlers
	r.GET("/redirect", redirectHandler) // 跳转到同意屏幕
	r.GET("/code", code)                // 接收授权服务器的code回调
	r.GET("/display", display)
	// 在5000端口启动
	r.Run(":5000")
}
