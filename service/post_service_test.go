package service

import (
	"testing"
	"goBlog/dao"
	"goBlog/model"
	"goBlog/config"
	
	"gorm.io/gorm"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
)

func setupPostServiceTestDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&model.User{}, &model.Post{})
	assert.NoError(t, err)
	config.DB = db
}

func TestCreatePost(t *testing.T) {
	setupPostServiceTestDB(t)

	// 创建用户
	user := &model.User{
		ID:       "u1",
		Account:  "acc001",
		Password: "passwd",
		Nickname: "nick",
		Avatar:   "avatar.png",
	}
	err := dao.CreateUser(user)
	assert.NoError(t, err)

	tests := []struct {
		name    string
		title   string
		content string
		cover   string
		userID  string
		wantErr bool
	}{
		{
			name:    "正常情况",
			title:   "测试标题",
			content: "这是内容",
			cover:   "cover.jpg",
			userID:  "u1",
			wantErr: false,
		},
		{
			name:    "空标题",
			title:   "",
			content: "有效内容",
			cover:   "cover.jpg",
			userID:  "u1",
			wantErr: true,
		},
		{
			name:    "空内容",
			title:   "有效标题",
			content: "",
			cover:   "cover.jpg",
			userID:  "u1",
			wantErr: true,
		},
		{
			name:    "空封面",
			title:   "标题",
			content: "内容",
			cover:   "",
			userID:  "u1",
			wantErr: true,
		},
		{
			name:    "用户不存在",
			title:   "标题",
			content: "内容",
			cover:   "封面.jpg",
			userID:  "non-existent",
			wantErr: true,
		},
		{
			name:    "标题超长",
			title:   "这是一段超过二十个字符的标题应该报错",
			content: "内容",
			cover:   "封面.jpg",
			userID:  "u1",
			wantErr: true,
		},
		{
			name:    "内容超长",
			title:   "标题",
			content: string(make([]byte, 301)),
			cover:   "封面.jpg",
			userID:  "u1",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			post, err := CreatePost(tt.title, tt.content, tt.cover, tt.userID)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, post)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, post)
				assert.Equal(t, tt.title, post.Title)
			}
		})
	}
}
