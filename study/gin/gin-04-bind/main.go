// Author:      xuan
// Date:        2023/6/21
// Description:	参数绑定

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Age      int    `form:"age" json:"age" binding:"required"`
}

func main() {
	engine := gin.Default()

	// 1.query  /user?username=seven&password=123456&age=24
	engine.GET("/user", func(context *gin.Context) {
		var user User
		// 也可以不指定绑定类型
		// err := context.ShouldBind(&user)
		err := context.ShouldBindWith(&user, binding.Query)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": err.Error()})
			return
		}
		fmt.Printf("%#v\n", user)
		context.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "data": user})
	})

	// 2.form
	engine.POST("/user/form", func(context *gin.Context) {
		var user User
		// _ = context.ShouldBind(&user)
		_ = context.ShouldBindWith(&user, binding.Form)
		fmt.Printf("%#v\n", user)
		context.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "data": user})
	})

	// 3.json
	engine.POST("/user/json", func(context *gin.Context) {
		var user User
		// _ = context.ShouldBind(&user)
		_ = context.ShouldBindJSON(&user)
		fmt.Printf("%#v\n", user)
		context.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "data": user})
	})

	err := engine.Run(":7777")
	if err != nil {
		return
	}
}
