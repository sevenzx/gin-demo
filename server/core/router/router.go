// Author:      xuan
// Date:        2023/6/25
// Description:	设置路由

package router

import (
	"gin-demo/server/core/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouters(engine *gin.Engine) {

	// 主页
	engine.GET("/", controller.HomePage)

	// 增删改查
	v1Group := engine.Group("/v1")
	{
		// 添加一个待办事项
		v1Group.POST("/todo", controller.SaveTodo)

		// 查看所有代办事项
		v1Group.GET("/todo", controller.ListTodo)

		// 更新一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateTodo)

		// 删除一个代办事项
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
}
