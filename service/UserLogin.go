package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
)

type UserLogin struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *UserLogin) Login() (model.User, serializer.Response) {
	var user model.User
	if err := model.DB.Where("account = ?", s.Account).Find(&user).Error; err != nil {
		return user, serializer.Response{
			Status: 1001,
			Msg:    "账号或密码错误",
			Error:  "账号或密码错误",
		}
	}
	PW, err := ScryptPw(s.Password)
	if err != nil {
		return user, serializer.Response{
			Status: 1001,
			Msg:    "账号或密码错误",
			Error:  "密码加密错误",
		}
	}
	if user.Password != PW {
		return user, serializer.Response{
			Status: 1001,
			Msg:    "账号或密码错误",
			Error:  "账号或密码错误",
		}
	}
	return user, serializer.Response{
		Status: 200,
		Msg:    "登录成功",
	}
}
