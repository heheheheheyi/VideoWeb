package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
)

type VideoInfo struct {
}

func (s *VideoInfo) GetInfo(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "获取数据失败",
			Error:  err.Error(),
		}
	}
	if video.Status == 1 {
		video.AddClick()
	}
	return serializer.BuildResponse(serializer.BuildVideo(video))
}
