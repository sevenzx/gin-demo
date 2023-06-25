// Author:      xuan
// Date:        2023/6/25
// Description:	Application主入口

package main

import (
	"fmt"
	"gin-demo/model/entity"
	"gin-demo/model/vo"
	"gin-demo/query"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

const dsn = "root:123456@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"

var (
	db *gorm.DB
)

func connectMySQL() (err error) {
	db, err = gorm.Open(mysql.Open(dsn))
	return err
}

func main() {

	// 连接数据库
	err := connectMySQL()
	if err != nil {
		fmt.Println("connect to MySQL failed, err ", err)
		return
	}

	// 设置数据库
	query.SetDefault(db)

	engine := gin.Default()
	// 加载HTML和静态文件
	engine.LoadHTMLGlob("templates/*.html")
	engine.Static("/static", "static")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 增删改查
	t := query.Todo
	v1Group := engine.Group("/v1")
	{
		// 添加一个待办事项
		v1Group.POST("/todo", func(c *gin.Context) {
			// 1.从请求中拿出数据
			var todo entity.Todo
			_ = c.BindJSON(&todo)
			// 2.插入数据库
			err = t.Create(&todo)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err})
			} else {
				c.JSON(http.StatusOK, gin.H{"msg": "ok"})
			}

		})

		// 查看所以代办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			// 查询数据库并数据脱敏
			todos, err2 := t.Find()
			if err2 != nil {
				c.JSON(http.StatusOK, gin.H{"error": err2})
			} else {
				var todoVOList = make([]vo.TodoVO, 0)
				for _, todo := range todos {
					var todoVO vo.TodoVO
					_ = copier.Copy(&todoVO, &todo)
					todoVOList = append(todoVOList, todoVO)
				}
				c.JSON(http.StatusOK, todoVOList)
			}
		})

		// 更新一个待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			// 1.判断id的合法性
			idStr := c.Param("id")
			id, err2 := strconv.ParseInt(idStr, 0, 64)
			if err2 != nil {
				c.JSON(http.StatusOK, gin.H{"error": "id不合法"})
			}
			// 2.从请求中拿出数据
			var todo entity.Todo
			_ = c.BindJSON(&todo)
			// 更新数据库
			_, err2 = t.Where(t.ID.Eq(id)).Update(t.Status, todo.Status)
			if err2 != nil {
				c.JSON(http.StatusOK, gin.H{"error": err2})
			} else {
				c.JSON(http.StatusOK, gin.H{"msg": "ok"})
			}
		})

		// 删除一个代办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			// 1.判断id的合法性
			idStr := c.Param("id")
			id, err2 := strconv.ParseInt(idStr, 0, 64)
			if err2 != nil {
				c.JSON(http.StatusOK, gin.H{"error": "id不合法"})
			}
			// 更新数据库
			_, err2 = t.Where(t.ID.Eq(id)).Delete()
			if err2 != nil {
				c.JSON(http.StatusOK, gin.H{"error": err2})
			} else {
				c.JSON(http.StatusOK, gin.H{"msg": "ok"})
			}
		})

	}

	err = engine.Run(":9000")
	if err != nil {
		return
	}
}
