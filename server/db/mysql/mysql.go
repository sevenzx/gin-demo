// Author:      xuan
// Date:        2023/6/25
// Description: MySQL配置

package mysql

import (
	"fmt"
	"gin-demo/server/config"
	"gin-demo/server/core/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL(database *config.MySQL) (err error) {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		database.Username,
		database.Password,
		database.Host,
		database.Port,
		database.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}

	// gen 设置数据库
	query.SetDefault(db)
	return nil
}
