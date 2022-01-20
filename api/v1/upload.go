package v1

import (
	"VideoWeb/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	var filesize = fileHeader.Size
	fmt.Println(filesize)
	if filesize >= 300000000 {
		c.JSON(http.StatusOK, gin.H{
			"status":  1000,
			"message": "文件过大",
			"error":   "文件过大",
		})
	}
	url, err := model.UpLoadFile(file, filesize)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1000,
			"message": "上传失败",
			"error":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "上传成功",
			"url":     url,
		})
	}
}
