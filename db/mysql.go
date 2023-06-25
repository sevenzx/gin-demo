// Author:      xuan
// Date:        2023/6/25
// Description: MySQL配置

package db

import (
	"gin-demo/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:123456@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectMySQL() (err error) {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}

	// gen 设置数据库
	query.SetDefault(db)
	return nil
}
