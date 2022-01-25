package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
	"fmt"
)

type VideoCreate struct {
	Title string `json:"title" binding:"required,max=50"`
	Info  string `json:"info" binding:"required,max=250"`
	URL   string `json:"url" binding:"required"`
	Img   string `json:"img" binding:"required"`
	Uid   uint
}

func (s *VideoCreate) Create() serializer.Response {
	fmt.Println(len(s.Title), len(s.Info))
	video := model.Video{
		Title:  s.Title,
		Info:   s.Info,
		URL:    s.URL,
		Img:    s.Img,
		Uid:    s.Uid,
		Status: 0,
	}
	if err := model.DB.Create(&video).Error; err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "投递视频失败",
			Error:  "数据库错误",
		}
	}
	video.AddClick()
	return serializer.Response{
		Status: 200,
		Msg:    "投递成功",
	}
}
