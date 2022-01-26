package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
)

type CommentAdminList struct {
}

func (s *CommentAdminList) GetList() serializer.Response {
	var comments []model.Comment
	total := 0

	if err := model.DB.Model(model.Comment{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "获取数据失败",
			Error:  err.Error(),
		}
	}
	err := model.DB.Find(&comments).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "获取数据失败",
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildComments(comments), uint(total))
}
