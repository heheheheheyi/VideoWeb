package service

import (
	"VideoWeb/cache"
	"VideoWeb/model"
	"VideoWeb/serializer"
	"strconv"
)

type VideoDelete struct {
}

func (s *VideoDelete) Delete(id string, uid uint) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}
	if video.Uid != uid {
		return serializer.Response{
			Status: 1000,
			Msg:    "无法删除视频",
			Error:  "无权限",
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
