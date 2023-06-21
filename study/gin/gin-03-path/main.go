// Author:      xuan
// Date:        2023/6/21
// Description:	网络请求的path参数

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	engine := gin.Default()

	engine.GET("/user/:name/:age", func(context *gin.Context) {
		name := context.Param("name")
		ageStr := context.Param("age")
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid age"})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	err := engine.Run(":7777")
	if err != nil {
		return
	}
}
