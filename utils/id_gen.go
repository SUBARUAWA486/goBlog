package utils

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// 生成 UUID 格式的 UserID/PostID
func GenerateUserID() string {
	return uuid.New().String()
}

func GeneratePostID() string {
	return uuid.New().String()
}

func GenerateStarsID() string {
	return uuid.New().String()
}

// 生成 12 位随机字符 CommentID
func GenerateCommentID() string {
	rand.Seed(time.Now().UnixNano())
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	id := ""
	for i := 0; i < 12; i++ {
		id += string(chars[rand.Intn(len(chars))])
	}
	return id
}
