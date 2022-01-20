package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
)

type UserRegister struct {
	Account   string `form:"account" json:"account" binding:"required,min=5,max=15"`
	Nickname  string `form:"nickname" json:"nickname"binding:"required,min=3,max=10"`
	Password  string `form:"password" json:"password"binding:"required,min=6,max=20"`
	ConfirmPW string `form:"confirm_pw" json:"confirm_pw"binding:"required,min=6,max=20"`
}

// Valid 表单验证
func (s *UserRegister) Valid() serializer.Response {
	if s.ConfirmPW != s.Password {
		return serializer.Response{
			Status: 1000,
			Msg:    "两次输入密码不一致",
			Error:  "两次输入密码不一致",
		}
	}
	count := 0
	model.DB.Model(&model.User{}).Where("account = ?", s.Account).Count(&count)
	if count > 0 {
		return serializer.Response{
			Status: 1000,
			Msg:    "账号已存在",
			Error:  "账号已存在",
		}
	}
	count = 0
	model.DB.Model(&model.User{}).Where("nickname = ?", s.Nickname).Count(&count)
	if count > 0 {
		return serializer.Response{
			Status: 1000,
			Msg:    "用户名已存在",
			Error:  "用户名已存在",
		}
	}
	return serializer.Response{
		Status: 200,
	}
}

// Register 用户注册
func (s *UserRegister) Register() serializer.Response {
	if response := s.Valid(); response.Status != 200 {
		return response
	}
	PW, err := ScryptPw(s.Password)
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "注册失败",
			Error:  err.Error(),
		}
	}
	userinfo := model.User{
		Account:  s.Account,
		Nickname: s.Nickname,
		Password: PW,
		Status:   1,
		Img:      "http://r50pyj4pr.hn-bkt.clouddn.com/account.png",
	}
	if err := model.DB.Create(&userinfo).Error; err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "注册失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "注册成功",
	}

}

func ScryptPw(password string) (string, error) {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{11, 22, 33, 44, 55, 66, 77, 88}
	HashPw, err := scrypt.Key([]byte(password), salt, 16, 8, 1, KeyLen)
	if err != nil {
		return "", err
	}
	FinalPw := base64.StdEncoding.EncodeToString(HashPw)
	return FinalPw, nil
}
