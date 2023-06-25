// Author:      xuan
// Date:        2023/6/24
// Description:	通过gen生成代码

package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

const dsn = "root:123456@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	g := gen.NewGenerator(gen.Config{
		// 指定实体类的路径 可省略
		ModelPkgPath: "./study/gen/model/entity",
		// 指定API路径
		OutPath: "./study/gen/query",
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
