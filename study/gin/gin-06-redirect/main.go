// Author:      xuan
// Date:        2023/6/21
// Description:	重定向

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	// 1.HTTP重定向
	engine.GET("/index", func(c *gin.Context) {
		// 301
		c.Redirect(http.StatusMovedPermanently, "https://wwww.baidu.com")
	})

	// 2.路由重定向
	engine.GET("/a", func(c *gin.Context) {
		// 这次的重定向不会改变地址栏的网址 还是 http://localhost:7777/a
		c.Request.URL.Path = "/b"
		engine.HandleContext(c)
	})
	engine.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "redirect ok"})
	})

	err := engine.Run(":7777")
	if err != nil {
		return
	}
}
