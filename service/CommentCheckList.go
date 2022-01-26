package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
)

type CommentCheckList struct {
}

func (s *CommentCheckList) GetList() serializer.Response {
	var comments []model.Comment
	total := 0

	if err := model.DB.Model(model.Comment{}).Where("status = 0").Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "获取数据失败",
			Error:  err.Error(),
		}
	}
	err := model.DB.Where("status = 0").Find(&comments).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "获取数据失败",
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildComments(comments), uint(total))
}
