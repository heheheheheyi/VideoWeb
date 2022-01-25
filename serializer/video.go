package serializer

import "VideoWeb/model"

// Video 视频序列化器
type Video struct {
	ID        uint    `json:"id"`
	Title     string  `json:"title"`
	Info      string  `json:"info"`
	URL       string  `json:"url"`
	Img       string  `json:"img"`
	Uid       uint    `json:"uid"`
	Status    int     `json:"status"`
	Donor     User    `json:"donor"`
	Comment   Comment `json:"comment"`
	Click     int     `json:"click"`
	CreatedAt string  `json:"created_at"`
}

// BuildVideo 序列化视频
func BuildVideo(video model.Video) Video {
	user, _ := model.GetUser(video.Uid)
	return Video{
		ID:        video.ID,
		Title:     video.Title,
		Info:      video.Info,
		URL:       video.URL,
		Img:       video.Img,
		Uid:       video.Uid,
		Status:    video.Status,
		Donor:     BuildUser(user),
		Click:     video.GetClick(),
		CreatedAt: video.CreatedAt.String()[:20],
	}
}

// BuildVideos 序列化视频列表
func BuildVideos(items []model.Video) (videos []Video) {
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}
