// Author:      xuan
// Date:        2023/6/25
// Description:	application主入口

package main

import (
	"fmt"
	"gin-demo/db"
	"gin-demo/route"
	"github.com/gin-gonic/gin"
)

func main() {

	// 连接数据库
	err := db.ConnectMySQL()
	if err != nil {
		fmt.Println("connect to MySQL failed, err ", err)
		return
	}

	// 初始化gin引擎
	engine := gin.Default()
	// 加载HTML和静态文件
	engine.LoadHTMLGlob("templates/*.html")
	engine.Static("/static", "static")

	// 设置路由
	route.SetupRouters(engine)

	err = engine.Run(":9000")
	if err != nil {
		return
	}
}
