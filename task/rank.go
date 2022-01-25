package task

import (
	"VideoWeb/cache"
	"VideoWeb/model"
)

func RestartDailyRank() error {
	return model.RedisClient.Del(cache.DailyRank).Err()
}
func RestartMonthlyRank() error {
	return model.RedisClient.Del(cache.MonthlyRank).Err()
}
func RestartDailyUpload() error {
	return model.RedisClient.Del(cache.DailyUpload).Err()

}
