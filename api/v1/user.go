package v1

import (
	"VideoWeb/middleware"
	"VideoWeb/model"
	"VideoWeb/serializer"
	"VideoWeb/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var data service.UserRegister
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, &serializer.Response{
			Status: 1000,
			Msg:    "数据格式错误",
			Error:  err.Error(),
		})
	} else {
		res := data.Register()
		c.JSON(http.StatusOK, res)
	}
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	var data service.UserLogin
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, &serializer.Response{
			Status: 1000,
			Msg:    "数据格式错误",
			Error:  err.Error(),
		})
	} else {
		user, response := data.Login()
		if response.Status != 200 {
			c.JSON(http.StatusOK, response)
		} else {
			token, err := middleware.GenerateToken(strconv.Itoa(int(user.ID)), user.Account)
			if err != nil {
				c.JSON(http.StatusOK, &serializer.Response{
					Status: 1000,
					Msg:    "请重新登录",
					Error:  err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, &serializer.Response{
					Status: 200,
					Data:   token,
				})
			}
		}
	}
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	c.JSON(http.StatusOK, serializer.Response{
		Status: 200,
		Msg:    "登出成功",
	})
}

// UserPersonal 用户个人信息
func UserPersonal(c *gin.Context) {
	id, _ := c.Get("user_id")
	user, err := model.GetUser(id)
	if err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			Status: 1000,
			Error:  err.Error(),
		})
	} else {
		res := serializer.BuildResponse(serializer.BuildUser(user))
		c.JSON(http.StatusOK, res)
	}
}

// UserUpdate 修改个人信息
func UserUpdate(c *gin.Context) {
	data := service.UserUpdate{}
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
		res := data.Update(user.ID)
		c.JSON(http.StatusOK, res)
	}
}

func PasswordUpdate(c *gin.Context) {
	data := service.PasswordUpdate{}
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
		res := data.Update(user.ID)
		c.JSON(http.StatusOK, res)
	}
}

func UserInfo(c *gin.Context) {
	var data service.UserInfo
	id := c.Param("id")
	res := data.GetUserInfo(id)
	c.JSON(http.StatusOK, res)
}
