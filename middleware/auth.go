package middleware

import (
	"VideoWeb/serializer"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// AuthRequired 验证登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenherder := c.Request.Header.Get("Authorization")
		if tokenherder == "" {
			c.JSON(200, serializer.Response{
				Status: 1000,
				Msg:    "无权限，请登录",
				Error:  "Token验证失败",
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenherder, " ", 2)
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.JSON(200, serializer.Response{
				Status: 1000,
				Msg:    "无权限，请登录",
				Error:  "Token验证失败",
			})
			c.Abort()
			return
		}
		key, Tcode := ParseToken(checkToken[1])
		if Tcode != nil {
			c.JSON(200, serializer.Response{
				Status: 1000,
				Msg:    "无权限，请登录",
				Error:  "Token验证失败",
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			c.JSON(200, serializer.Response{
				Status: 1000,
				Msg:    "无权限，请登录",
				Error:  "Token验证失败",
			})
			c.Abort()
			return
		}
		c.Set("user_id", key.Id)
		c.Next()
	}
}
