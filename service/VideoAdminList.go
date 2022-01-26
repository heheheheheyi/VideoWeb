package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
)

type VideoAdminList struct {
}

func (s *VideoAdminList) GetList() serializer.Response {
	var videos []model.Video
	total := 0

	if err := model.DB.Model(model.Video{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "获取数据失败",
			Error:  err.Error(),
		}
	}
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "获取数据失败",
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildVideos(videos), uint(total))
}
