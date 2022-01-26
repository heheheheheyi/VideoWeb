package v1

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
	"VideoWeb/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func VideoCreate(c *gin.Context) {
	var data service.VideoCreate
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusOK, &serializer.Response{
			Status: 1000,
			Msg:    "数据格式错误",
			Error:  err.Error(),
		})
	} else {
		id, _ := c.Get("user_id")
		user, _ := model.GetUser(id)
		if user.Nickname == "" {
			c.JSON(http.StatusOK, &serializer.Response{
				Status: 1000,
				Msg:    "无权限，请登录",
				Error:  "jwt错误",
			})
		} else {
			data.Uid = user.ID
			res := data.Create()
			c.JSON(http.StatusOK, res)
		}
	}
}
func VideoUpdate(c *gin.Context) {
	data := service.VideoUpdate{}
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			Status: 1000,
			Msg:    "输入数据有误",
			Error:  err.Error(),
		})
	} else {
		id, _ := c.Get("user_id")
		user, _ := model.GetUser(id)
		res := data.Update(c.Param("id"), user.ID)
		c.JSON(http.StatusOK, res)
	}
}
func VideoDelete(c *gin.Context) {
	data := service.VideoDelete{}
	id, _ := c.Get("user_id")
	user, _ := model.GetUser(id)
	res := data.Delete(c.Param("id"), user.ID)
	c.JSON(http.StatusOK, res)
}
func VideoList(c *gin.Context) {
	var s service.VideoList
	s.PageSize, _ = strconv.Atoi(c.Query("pagesize"))
	s.PageNum, _ = strconv.Atoi(c.Query("pagenum"))
	s.Desc, _ = strconv.Atoi(c.Query("desc"))
	s.OrderElement = c.Query("orderelement")
	c.JSON(http.StatusOK, s.GetList())
}

func UserVideoList(c *gin.Context) {
	var data service.UserVideoList
	data.Uid, _ = strconv.Atoi(c.Param("uid"))
	c.JSON(http.StatusOK, data.GetList())
}
func VideoInfo(c *gin.Context) {
	var data service.VideoInfo
	res := data.GetInfo(c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func VideoCheckList(c *gin.Context) {
	var data service.VideoCheckList
	c.JSON(http.StatusOK, data.GetList())
}
