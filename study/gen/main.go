// Author:      xuan
// Date:        2023/6/24
// Description:	操作数据库

package main

import (
	"fmt"
	"gin-demo/study/gen/dal/model/gender"
	"gin-demo/study/gen/dal/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:123456@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db, _ := gorm.Open(mysql.Open(dsn))
	// 设置默认数据库
	query.SetDefault(db)
	u := query.User

	// 软删除测试成功
	// info, _ := u.Where(u.ID.Eq(1)).Delete()
	// fmt.Println(info)

	users, _ := u.Find()
	for _, user := range users {
		fmt.Println(user)
	}

	// 多条件查询
	user, _ := u.Where(u.Gender.Eq(gender.Male), u.Username.Eq("mars")).First()
	fmt.Println(user)

	// 更新
	_, _ = u.Where(u.ID.Eq(2)).Update(u.Username, "Ease")
}
