package model

import (
	"VideoWeb/cache"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Video struct {
	gorm.Model
	Title  string `gorm:"not null"`
	Info   string `gorm:"not null"`
	URL    string `gorm:"not null"`
	Uid    uint   `gorm:"not null"`
	Img    string `gorm:"not null"`
	Status int    `gorm:"not null;default:0;index:index_status"`
}

func ClickKey(id uint) string {
	return fmt.Sprintf("click:video:%s", strconv.Itoa(int(id)))
}
func (video *Video) GetClick() int {
	str, _ := RedisClient.Get(ClickKey(video.ID)).Result()
	count, _ := strconv.Atoi(str)
	return count
}
func (video *Video) AddClick() {
	RedisClient.Incr(ClickKey(video.ID))
	RedisClient.ZIncrBy(cache.Clickrank, 1, strconv.Itoa(int(video.ID)))
	RedisClient.ZIncrBy(cache.DailyRank, 1, strconv.Itoa(int(video.ID)))
	RedisClient.ZIncrBy(cache.MonthlyRank, 1, strconv.Itoa(int(video.ID)))
}
