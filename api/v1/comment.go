package v1

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
	"VideoWeb/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CommentCreate(c *gin.Context) {
	data := service.CommentCreate{}
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
		vid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusOK, serializer.Response{
				Status: 1000,
				Msg:    "评论失败",
				Error:  err.Error(),
			})
		} else {
			res := data.Create(uint(vid), user.ID)
			c.JSON(http.StatusOK, res)
		}
	}
}

func CommentDelete(c *gin.Context) {
	data := service.CommentDelete{}
	id, _ := c.Get("user_id")
	user, _ := model.GetUser(id)
	res := data.Delete(c.Param("id"), user.ID)
	c.JSON(http.StatusOK, res)
}

func CommentList(c *gin.Context) {
	data := service.CommentList{}
	data.PageSize, _ = strconv.Atoi(c.Query("pagesize"))
	data.PageNum, _ = strconv.Atoi(c.Query("pagenum"))
	data.Desc, _ = strconv.Atoi(c.Query("desc"))
	data.OrderElement = c.Query("orderelement")
	c.JSON(http.StatusOK, data.GetList(c.Param("id")))
}

func CommentCheckList(c *gin.Context) {
	var data service.CommentCheckList
	c.JSON(http.StatusOK, data.GetList())
}

func CommentCheckPass(c *gin.Context) {
	var data service.CommentCheckPass
	c.JSON(http.StatusOK, data.Pass(c.Param("id")))
}

func CommentCheckBan(c *gin.Context) {
	var data service.CommentCheckBan
	c.JSON(http.StatusOK, data.Ban(c.Param("id")))
}
func CommentAdminList(c *gin.Context) {
	var s service.CommentAdminList
	c.JSON(http.StatusOK, s.GetList())
}
func CommentAdminDelete(c *gin.Context) {
	data := service.CommentAdminDelete{}
	res := data.Delete(c.Param("id"))
	c.JSON(http.StatusOK, res)
}
