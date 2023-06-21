// Author:      xuan
// Date:        2023/6/21
// Description:	文件上传

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLFiles("study/gin/gin-05-upload/index.html")

	engine.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// engine.MaxMultipartMemory = 8 << 20  // 8 MiB

	// 上传单个文件
	//engine.POST("/upload", func(context *gin.Context) {
	//	// 与 <input type="file" name="file"> name对应
	//	file, err := context.FormFile("file")
	//	if err != nil {
	//		context.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": err.Error()})
	//		return
	//	}
	//
	//	dst := path.Join("./file/", file.Filename)
	//	_ = context.SaveUploadedFile(file, dst)
	//	context.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok"})
	//})

	// 上传多个文件
	engine.POST("/upload", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		files := form.File["file"]

		for _, file := range files {
			dst := path.Join("./file/", file.Filename)
			_ = context.SaveUploadedFile(file, dst)
		}
		context.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok"})
	})

	err := engine.Run(":7777")
	if err != nil {
		return
	}
}
