package service

import (
	"VideoWeb/cache"
	"VideoWeb/model"
	"VideoWeb/serializer"
	"strconv"
)

type VideoAdminDelete struct {
}

func (s *VideoAdminDelete) Delete(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "视频删除失败",
			Error:  err.Error(),
		}
	}
	model.RedisClient.ZRem(cache.Clickrank, strconv.Itoa(int(video.ID)))
	model.RedisClient.ZRem(cache.MonthlyRank, strconv.Itoa(int(video.ID)))
	model.RedisClient.ZRem(cache.DailyRank, strconv.Itoa(int(video.ID)))
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}
