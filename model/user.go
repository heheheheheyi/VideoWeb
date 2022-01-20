package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Account  string
	Nickname string
	Password string
	Img      string
	Status   uint
}
type UserInfo struct {
	Id       uint
	Account  string
	Nickname string
	Img      string
}

func GetUser(id interface{}) (User, error) {
	var user User
	res := DB.First(&user, id)
	return user, res.Error
}
