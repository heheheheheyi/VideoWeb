package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
)

type PasswordUpdate struct {
	Password    string `form:"password" json:"password"binding:"required,min=6,max=20"`
	NewPassword string `form:"new_password" json:"new_password" binding:"required,min=6,max=20"`
	ConfirmPW   string `form:"confirm_pw" json:"confirm_pw"binding:"required,min=6,max=20"`
}

func (s *PasswordUpdate) Update(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	PW, _ := ScryptPw(s.Password)
	if user.Password != PW {
		return serializer.Response{
			Status: 1000,
			Msg:    "密码错误",
		}
	}
	if s.ConfirmPW != s.NewPassword {
		return serializer.Response{
			Status: 1000,
			Msg:    "两次输入密码不一致",
			Error:  "两次输入密码不一致",
		}
	}
	user.Password, _ = ScryptPw(s.NewPassword)
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
