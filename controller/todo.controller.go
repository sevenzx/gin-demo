// Author:      xuan
// Date:        2023/6/25
// Description: controller层

package controller

import (
	"gin-demo/model/entity"
	"gin-demo/model/vo"
	"gin-demo/query"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"strconv"
)

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func SaveTodo(c *gin.Context) {
	var t = query.Todo
	// 1.从请求中拿出数据
	var todo entity.Todo
	_ = c.BindJSON(&todo)
	// 2.插入数据库
	err := t.Create(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	}
}

func ListTodo(c *gin.Context) {
	var t = query.Todo
	// 查询数据库并数据脱敏
	todos, err := t.Find()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		var todoVOList = make([]vo.TodoVO, 0)
		for _, todo := range todos {
			var todoVO vo.TodoVO
			_ = copier.Copy(&todoVO, &todo)
			todoVOList = append(todoVOList, todoVO)
		}
		c.JSON(http.StatusOK, todoVOList)
	}
}

func UpdateTodo(c *gin.Context) {
	var t = query.Todo
	// 1.判断id的合法性
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "id不合法"})
	}
	// 2.从请求中拿出数据
	var todo entity.Todo
	_ = c.BindJSON(&todo)
	// 更新数据库
	_, err = t.Where(t.ID.Eq(id)).Update(t.Status, todo.Status)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	}
}

func DeleteTodo(c *gin.Context) {
	var t = query.Todo
	// 1.判断id的合法性
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "id不合法"})
	}
	// 更新数据库
	_, err = t.Where(t.ID.Eq(id)).Delete()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	}
}
