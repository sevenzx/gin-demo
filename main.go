// Author:      xuan
// Date:        2023/6/25
// Description:	Application主入口

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	// 加载HTML和静态文件
	engine.LoadHTMLGlob("templates/*.html")
	engine.Static("/static", "static")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	err := engine.Run(":9000")
	if err != nil {
		return
	}
}
