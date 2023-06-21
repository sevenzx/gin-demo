// Author: 		xuan
// Date: 		2023/6/21
// Description: 获取form参数

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	// Gin框架中使用LoadHTMLGlob()或者LoadHTMLFiles()方法进行HTML模板渲染
	// Gin会将该文件注册为模板，模板名称为"login.html" "home.html"
	engine.LoadHTMLFiles("study/gin/gin-02-form/login.html", "study/gin/gin-02-form/home.html")

	engine.GET("/login", func(context *gin.Context) {
		// name 需要使用模板名称
		context.HTML(http.StatusOK, "login.html", nil)
	})

	engine.POST("/login", func(context *gin.Context) {
		// 1.基础的获取
		//username := context.PostForm("username")
		//password := context.PostForm("password")
		// 2.判断是否存在
		username, ok := context.GetPostForm("username")
		if !ok {
			username = "somebody"
		}
		password, _ := context.GetPostForm("password")
		context.HTML(http.StatusOK, "home.html", gin.H{
			"Name": username,
			"Pwd":  password,
		})
	})

	err := engine.Run(":7777")
	if err != nil {
		return
	}
}
