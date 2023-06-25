// Author:      xuan
// Date:        2023/6/25
// Description:	通过gen生成代码

package main

import (
	"fmt"
	"gin-demo/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

var dsn string

func setDsn() {
	err := setting.InitSetting()
	if err != nil {
		fmt.Printf("Unable to parse the configuration file: %v\n", err)
		return
	}
	database := setting.Config.Database
	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		database.Username,
		database.Password,
		database.Host,
		database.Port,
		database.DB,
	)
}

func main() {
	setDsn()
	g := gen.NewGenerator(gen.Config{
		// 指定实体类的路径 可省略
		ModelPkgPath: "./model/entity",
		// 指定API路径
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	db, _ := gorm.Open(mysql.Open(dsn))
	g.UseDB(db) // reuse your gorm db

	// 为所有的数据库表使用下列的配置
	g.ApplyBasic(g.GenerateAllTable(
		// 将数据库中is_delete的字段的type中对应的类型设置为soft_delete.DeletedAt
		// 然后soft_delete.DeletedAt类型就会让gorm知道该表开启了软删除(逻辑删除)
		gen.FieldType("is_delete", "soft_delete.DeletedAt"),
		gen.FieldGORMTag("is_delete", func(tag field.GormTag) field.GormTag {
			// 通过tag的方式将softDelete设置为flag模式(0/1)
			tag.Set("softDelete", "flag")
			return tag
		}),
	)...)

	// Generate the code
	g.Execute()
}
