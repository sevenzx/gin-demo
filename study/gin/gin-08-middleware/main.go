// Author:      xuan
// Date:        2023/6/22
// Description:	中间件

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 计算耗时中间件
func timeMiddleware(context *gin.Context) {
	fmt.Println("time in...")
	start := time.Now()
	// 调用后续的处理函数
	context.Next()
	cost := time.Since(start)
	fmt.Println("cost : ", cost)
	fmt.Println("time out...")
}

// 通过这个中间件来了解
// context.Next()的执行顺序、context.Abort()的用法
// 跨中间件传递值、中间件起goroutine
func testMiddleware(c *gin.Context) {
	fmt.Println("test in..")
	// 可以通过c.Get("key") 获取值
	c.Set("middlewareMsg", "middle")
	//go func(c *gin.Context) {
	//	// 这里只能使用c.Copy() 只能读取不能修改
	//}(c.Copy())
	c.Next()
	//c.Abort() // 阻值调用后续的处理函数
	fmt.Println("test out...")
}

// 中间件的一般写法
func authMiddleware(doCheck bool) gin.HandlerFunc {
	// 查询数据库...
	// 其他操作等
	return func(c *gin.Context) {
		if doCheck {
			// 一些业务逻辑
			c.Next()
		} else {
			c.Next()
		}
	}
}

func main() {
	// gin.Default()使用了Logger Recovery中间件
	// 源码: 	engine.Use(Logger(), Recovery())
	// 想使用一个没有任何中间件的engin 可以使用gin.New()
	engine := gin.Default()
	// 	全局注册中间件
	engine.Use(timeMiddleware, testMiddleware)

	// 某个请求单独使用中间件
	engine.GET("/index", timeMiddleware, func(context *gin.Context) {
		fmt.Println("/index")
		// 从上下文中取值
		msg, _ := context.Get("middlewareMsg")
		context.JSON(http.StatusOK, gin.H{"code": 0, "msg": msg})
	})

	// 路由组使用中间件有两种写法
	// 1.
	xxGroup := engine.Group("/xx", timeMiddleware)
	// 2.
	xxGroup.Use(authMiddleware(true))

	err := engine.Run(":7777")
	if err != nil {
		return
	}
}
