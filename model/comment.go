package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Info   string
	Vid    uint
	Uid    uint
	Status int `gorm:"not null;default:0;index:index_status"`
}
