package service

import (
	"VideoWeb/model"
	"VideoWeb/serializer"
	"fmt"
)

type CommentList struct {
	PageNum      int
	PageSize     int
	OrderElement string
	Desc         int
}

func (s *CommentList) GetList(id string) serializer.Response {
	var Comments []model.Comment
	total := 0

	if err := model.DB.Model(model.Comment{}).Where("status = 1 and vid = ?", id).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 1000,
			Msg:    "获取数据失败",
			Error:  err.Error(),
		}
	}
	if s.Desc == 1 {
		OrderElement := fmt.Sprint(s.OrderElement, " desc")
		err := model.DB.Where("status = 1 and vid = ?", id).Limit(s.PageSize).Offset((s.PageNum - 1) * s.PageSize).Order(OrderElement).Find(&Comments).Error
		if err != nil {
			return serializer.Response{
				Status: 1000,
				Msg:    "获取数据失败",
				Error:  err.Error(),
			}
		}
	} else {
		err := model.DB.Where("status = 1 and vid = ?", id).Limit(s.PageSize).Offset((s.PageNum - 1) * s.PageSize).Order(s.OrderElement).Find(&Comments).Error
		if err != nil {
			return serializer.Response{
				Status: 1000,
				Msg:    "获取数据失败",
				Error:  err.Error(),
			}
		}
	}
	return serializer.BuildListResponse(serializer.BuildComments(Comments), uint(total))
}
