package middleware

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"strings"
)

func LimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//得到ip地址
		ipAddr := GetRealIp(c)
		fmt.Println("current ip:" + ipAddr)
		//ipAddr:="127.0.0.1"
		limiter := model.RateLimiter.GetLimiter(ipAddr)
		if !limiter.Allow() {
			c.JSON(200, serializer.Response{
				Status: 1000,
				Msg:    "访问过快",
				Error:  "访问过快",
			})
			c.Abort()
		} else {
			c.Next()
		}
	}
}
func GetRealIp(r *gin.Context) string {
	ip, _, err := net.SplitHostPort(r.Request.RemoteAddr)
	if err != nil {
		ip = r.Request.RemoteAddr
	}
	if ip != "127.0.0.1" {
		return ip
	}
	// Check if behide nginx or apache
	xRealIP := r.Request.Header.Get("X-Real-Ip")
	xForwardedFor := r.Request.Header.Get("X-Forwarded-For")

	for _, address := range strings.Split(xForwardedFor, ",") {
		address = strings.TrimSpace(address)
		if address != "" {
			return address
		}
	}

	if xRealIP != "" {
		return xRealIP
	}
	return ip
}
