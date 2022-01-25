package router

import (
	"VideoWeb/api/v1"
	"VideoWeb/config"
	"VideoWeb/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {

	gin.SetMode(config.GinMode)

	r := gin.Default()

	r.Use(middleware.Cors()).Use(middleware.Logger()).Use(middleware.LimitMiddleware())

	router := r.Group("api/v1")
	{
		//用户注册
		router.POST("user/register", v1.UserRegister)

		//用户登录
		router.POST("user/login", v1.UserLogin)

		//用户信息
		router.GET("user/:id", v1.UserInfo)

		//视频列表
		router.GET("video", v1.VideoList)

		//单个视频
		router.GET("video/:id", v1.VideoInfo)

		//视频评论列表
		router.GET("comment/:id", v1.CommentList)

		//日排行榜
		router.GET("dailyrank", v1.DailyRank)

		//月排行榜
		router.GET("monthlyrank", v1.MonthlyRank)

		//总排行
		router.GET("clickrank", v1.ClickRank)
	}

	// 需要登录
	authed := router.Group("/")
	authed.Use(middleware.AuthRequired())
	{
		//个人信息
		authed.GET("user/personal", v1.UserPersonal)

		//登出
		authed.DELETE("user/logout", v1.UserLogout)

		//修改个人信息
		authed.PUT("userinfo", v1.UserUpdate)

		//修改密码
		authed.PUT("password", v1.PasswordUpdate)

		//添加视频
		authed.POST("video", v1.VideoCreate)

		//更新视频
		authed.PUT("video/:id", v1.VideoUpdate)

		//用戶投稿視頻
		router.GET("uservideolist/:uid", v1.UserVideoList)

		//删除视频
		authed.DELETE("video/:id", v1.VideoDelete)

		//添加评论
		authed.POST("comment/:id", v1.CommentCreate)

		//删除评论
		authed.DELETE("comment/:id", v1.CommentDelete)

		//上传七牛云
		authed.POST("upload", v1.UpLoad)
	}

	r.Run(config.GinPort)
}
