package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	err := engine.Run(":7777")
	if err != nil {
		return
	}

}
