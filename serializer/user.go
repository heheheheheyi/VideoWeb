package serializer

import "VideoWeb/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	Account   string `json:"account"`
	Nickname  string `json:"nickname"`
	Status    int    `json:"status"`
	Img       string `json:"img"`
	CreatedAt string `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		Account:   user.Account,
		Nickname:  user.Nickname,
		Img:       user.Img,
		Status:    user.Status,
		CreatedAt: user.CreatedAt.String()[:20],
	}
}
