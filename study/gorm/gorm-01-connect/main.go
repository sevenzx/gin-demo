// Author:      xuan
// Date:        2023/6/23
// Description:	连接MySQL数据库

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID     int64 `gorm:"primaryKey"`
	Name   string
	Age    int
	Gender string
	Status int `gorm:"default:1"`
}

func (User) TableName() string {
	return "user"
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn))

	// 1.创建表 自动迁移(将结构体和数据表进行同步)
	_ = db.AutoMigrate(&User{})

	// 2.插入数据
	//db.Create(&User{Name: "Mary", Age: 24, Gender: "女"})

	// 3.查找数据
	var user = new(User)
	//db.Where("name = ?", "Mars").First(&user)
	db.First(&user, "name = ?", "Jack")
	fmt.Printf("%#v\n", user)

	// 3.修改数据
	db.Model(&user).Update("age", 24)

	// 4.删除数据
	db.Delete(&user)
}
