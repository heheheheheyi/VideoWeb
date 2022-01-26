package middleware

import (
	"VideoWeb/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Admin 管理员登录
func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		str, _ := c.Get("user_id")
		id, _ := strconv.Atoi(str.(string))
		if id != 1 {
			c.JSON(200, serializer.Response{
				Status: 1000,
				Msg:    "无管理员权限",
				Error:  "无管理员权限",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
