package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
)

type UserInfo struct {
	UserInfo model.UserInfo
}

func (s *UserInfo) GetUserInfo(id string) serializer.Response {
	var user model.User
	err := model.DB.First(&user, id).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "数据获取失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildUser(user),
	}
}
