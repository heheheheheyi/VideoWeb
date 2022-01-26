package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
)

type CommentAdminDelete struct {
}

func (s *CommentAdminDelete) Delete(id string) serializer.Response {
	var comment model.Comment
	err := model.DB.First(&comment, id).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&comment).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "视频删除失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}
