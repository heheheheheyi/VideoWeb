package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
	"fmt"
)

type VideoList struct {
	PageNum      int
	PageSize     int
	OrderElement string
	Desc         int
}

func (s *VideoList) GetList() serializer.Response {
	var videos []model.Video
	total := 0

	if err := model.DB.Model(model.Video{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "获取数据失败",
			Error:  err.Error(),
		}
	}
	if s.Desc == 1 {
		OrderElement := fmt.Sprint(s.OrderElement, " desc")
		err := model.DB.Limit(s.PageSize).Offset((s.PageNum - 1) * s.PageSize).Order(OrderElement).Find(&videos).Error
		if err != nil {
			return serializer.Response{
				Status: 1000,
				Msg:    "获取数据失败",
				Error:  err.Error(),
			}
		}
	} else {
		err := model.DB.Limit(s.PageSize).Offset((s.PageNum - 1) * s.PageSize).Order(s.OrderElement).Find(&videos).Error
		if err != nil {
			return serializer.Response{
				Status: 1000,
				Msg:    "获取数据失败",
				Error:  err.Error(),
			}
		}
	}
	return serializer.BuildListResponse(serializer.BuildVideos(videos), uint(total))
}
