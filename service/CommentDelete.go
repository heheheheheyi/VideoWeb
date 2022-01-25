package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
)

type CommentDelete struct {
}

func (s *CommentDelete) Delete(id string, uid uint) serializer.Response {
	var comment model.Comment
	err := model.DB.First(&comment, id).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}
	if comment.Uid != uid {
		return serializer.Response{
			Status: 1000,
			Msg:    "无权限删除该评论",
			Error:  "无权限",
		}
	}
	err = model.DB.Delete(&comment).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "删除评论失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除评论成功",
	}
}
