package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
)

type UserVideoList struct {
	Uid int
}

func (s *UserVideoList) GetList() serializer.Response {
	var videos []model.Video
	total := 0

	if err := model.DB.Model(model.Video{}).Where("uid = ?", s.Uid).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "获取数据失败",
			Error:  err.Error(),
		}
	}
	err := model.DB.Where("uid = ?", s.Uid).Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "获取数据失败",
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildVideos(videos), uint(total))
}
