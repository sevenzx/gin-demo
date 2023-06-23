// Author:      xuan
// Date:        2023/6/23
// Description:	模型和表结构 查看官网:https://gorm.io/zh_CN/docs/models.html

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Blog struct {
	gorm.Model
	ID          int64
	Name        string
	Description string
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 开启单数表名
			SingularTable: true,
			// 表前缀
			TablePrefix: "tb_",
		},
	})

	_ = db.AutoMigrate(&Blog{})
	//blog := Blog{
	//	Name:        "如何提高编程能力",
	//	Description: "熬夜学习...",
	//}
	//
	//db.Create(&blog)
	var blog Blog
	db.Find(&blog, 1)
	fmt.Println(blog)
}
