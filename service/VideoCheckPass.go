package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
)

type VideoCheckPass struct {
}

func (s *VideoCheckPass) Pass(id string) serializer.Response {
	var video model.Video
	model.DB.First(&video, id)
	video.Status = 1
	err := model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "审核失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "审核成功",
	}
}
