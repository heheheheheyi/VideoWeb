package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
)

type CommentCreate struct {
	Info string `json:"info" form:"info" binding:"required,max=250"`
}

func (s *CommentCreate) Create(vid uint, uid uint) serializer.Response {
	var comment model.Comment
	comment.Uid = uid
	comment.Vid = vid
	comment.Info = s.Info
	if err := model.DB.Create(&comment).Error; err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "创建视频失败",
			Error:  "数据库错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "评论成功",
	}
}
