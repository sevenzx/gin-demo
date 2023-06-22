// Author:      xuan
// Date:        2023/6/22
// Description:	路由

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	// 1.普通路由 指定请求方式
	engine.GET("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok"})
	})

	// 2.Any 处理任何请求
	engine.Any("/any", func(context *gin.Context) {
		//anyMethods = []string{  源码中的内容
		//	http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch,
		//	http.MethodHead, http.MethodOptions, http.MethodDelete, http.MethodConnect,
		//	http.MethodTrace,
		//}
		context.JSON(http.StatusOK, gin.H{"code": 0, "msg": "any"})
	})

	// 3.NoRoute 处理没有匹配到路由的请求
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"code": -1, "msg": "not found"})
	})

	// 4.路由组
	userGroup := engine.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "/user/index"})
		})
		userGroup.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "/user/login"})
		})

		// 路由组里还能再次嵌套
		v1Group := userGroup.Group("/v1")
		{
			v1Group.GET("/xx", func(context *gin.Context) {
			})
		}
	}

	err := engine.Run(":7777")
	if err != nil {
		return
	}
}
