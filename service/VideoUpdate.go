package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
)

type VideoUpdate struct {
	Title string `json:"title" binding:"required,max=50"`
	Info  string `json:"info" binding:"required,max=250"`
	Img   string `json:"img" binding:"required"`
}

func (s *VideoUpdate) Update(id string, uid uint) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}
	if video.Uid != uid {
		return serializer.Response{
			Status: 1000,
			Msg:    "无权限更新",
			Error:  "无权限",
		}
	}
	video.Title = s.Title
	video.Info = s.Info
	video.Img = s.Img
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "更新失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "更新成功",
	}
}
