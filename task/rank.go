package task

import "VideoWeb/model"

func RestartDailyRank() error {
	return model.RedisClient.Del("rank:daily").Err()
}
func RestartMonthlyRank() error {
	return model.RedisClient.Del("rank:monthly").Err()
}
