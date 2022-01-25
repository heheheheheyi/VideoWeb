package serializer

import "VideoWeb/model"

type Comment struct {
	ID        uint   `json:"id"`
	Info      string `json:"info"`
	Vid       uint   `json:"vid"`
	Uid       uint   `json:"uid"`
	User      User   `json:"donor"`
	CreatedAt string `json:"created_at"`
}

// BuildComment 序列化视频
func BuildComment(comment model.Comment) Comment {
	user, _ := model.GetUser(comment.Uid)
	return Comment{
		ID:        comment.ID,
		Info:      comment.Info,
		Vid:       comment.Vid,
		Uid:       comment.Uid,
		User:      BuildUser(user),
		CreatedAt: comment.CreatedAt.String()[:20],
	}
}

// BuildComments 序列化视频列表
func BuildComments(items []model.Comment) (comments []Comment) {
	for _, item := range items {
		comment := BuildComment(item)
		comments = append(comments, comment)
	}
	return comments
}
