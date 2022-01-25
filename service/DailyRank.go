package service

import (
	"VideoWeb/cache"
	"VideoWeb/model"
	"VideoWeb/serializer"
	"fmt"
	"strings"
)

type DailyRank struct {
}

func (s *DailyRank) GetRank() serializer.Response {
	var videos []model.Video

	ids, _ := model.RedisClient.ZRevRange(cache.DailyRank, 0, 9).Result()

	if len(ids) > 0 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(ids, ","))
		err := model.DB.Where("status = 1 and id in (?)", ids).Order(order).Find(&videos).Error
		if err != nil {
			return serializer.Response{
				Status: 1000,
				Msg:    "数据库错误",
				Error:  err.Error(),
			}
		}
	}

	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildVideos(videos),
	}
}
