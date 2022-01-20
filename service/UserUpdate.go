package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
	"fmt"
)

type UserUpdate struct {
	Nickname string `form:"nickname" json:"nickname"binding:"required,min=3,max=10"`
	Img      string `json:"img" binding:"required"`
}

func (s *UserUpdate) Update(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	if user.Nickname != s.Nickname {
		count := 0
		model.DB.Model(&model.User{}).Where("nickname = ?", s.Nickname).Count(&count)
		if count > 0 {
			return serializer.Response{
				Status: 1000,
				Msg:    "用户名已存在",
				Error:  "用户名已存在",
			}
		}
	}
	user.Nickname = s.Nickname
	user.Img = s.Img
	fmt.Println(s.Img)
	err := model.DB.Save(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "更新失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "更新成功",
	}
}
