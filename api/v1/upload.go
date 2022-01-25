package v1

import (
	"VideoWeb/cache"
	"VideoWeb/model"
	"VideoWeb/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpLoad(c *gin.Context) {
	model.RedisClient.Incr(cache.DailyUpload)
	str, _ := model.RedisClient.Get(cache.DailyUpload).Result()
	count, _ := strconv.Atoi(str)
	if count >= 20 {
		c.JSON(http.StatusOK, serializer.Response{
			Status: 1000,
			Msg:    "上传数量上限,请明日再试",
		})
		return
	}
	file, fileHeader, _ := c.Request.FormFile("file")
	var filesize = fileHeader.Size
	fmt.Println(filesize)
	if filesize >= 300000000 {
		c.JSON(http.StatusOK, serializer.Response{
			Status: 1000,
			Msg:    "文件过大",
			Error:  "文件过大",
		})
		return
	}
	url, err := model.UpLoadFile(file, filesize)
	if err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			Status: 1000,
			Msg:    "上传失败",
			Error:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "上传成功",
			"url":    url,
		})
	}
}
