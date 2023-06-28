// Author:      xuan
// Date:        2023/6/25
// Description:	application主入口

package main

import (
	"fmt"
	"gin-demo/server/config"
	"gin-demo/server/core/router"
	"gin-demo/server/db/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	// 从配置文件中加载配置
	err := config.InitSetting()
	if err != nil {
		return
	}

	// 连接数据库
	err = mysql.ConnectMySQL(&config.Config.MySQL)
	if err != nil {
		fmt.Println("connect to MySQL failed, err ", err)
		return
	}

	// 初始化gin引擎
	engine := gin.Default()
	// 加载HTML和静态文件
	engine.LoadHTMLGlob("web/templates/*.html")
	engine.Static("/static", "web/static")

	// 设置路由
	router.SetupRouters(engine)

	err = engine.Run(fmt.Sprintf(":%d", config.Config.Port))
	if err != nil {
		return
	}
}
