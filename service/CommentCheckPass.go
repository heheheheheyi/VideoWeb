package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
)

type CommentCheckPass struct {
}

func (s *CommentCheckPass) Pass(id string) serializer.Response {
	var comment model.Comment
	model.DB.First(&comment, id)
	comment.Status = 1
	err := model.DB.Save(&comment).Error
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
