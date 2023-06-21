// Author: 		xuan
// Date: 		2023/6/21
// Description: 获取query string参数

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Day struct {
	Date    int    `json:"date"`
	DateStr string `json:"dateStr"`
	Year    int    `json:"year"`
	YearDay int    `json:"yearDay"`
	Week    int    `json:"week"`
	Weekend int    `json:"weekend"`
	Workday int    `json:"workday"`
	Holiday string `json:"holiday"`
}

func main() {
	engine := gin.Default()

	// 普通的get请求
	engine.GET("/ping", func(context *gin.Context) {
		day := Day{
			Date:    20230614,
			DateStr: "2023-06-14",
			Year:    2023,
			YearDay: 165,
			Week:    3,
			Weekend: 0,
			Workday: 1,
			Holiday: "非节假日",
		}
		context.JSON(http.StatusOK, day)
	})

	// 获取query string
	engine.GET("/search", func(context *gin.Context) {
		// 1.基础的获取
		name := context.Query("name")
		// 2.如果没有传递参数会使用默认值
		age := context.DefaultQuery("age", "18")
		// 3.判断是否有值
		hobby, ok := context.GetQuery("hobby")
		if !ok {
			hobby = "nothing"
		}
		context.JSON(http.StatusOK, gin.H{"name": name, "age": age, "hobby": hobby})
	})

	err := engine.Run(":7777")
	if err != nil {
		return
	}

}
