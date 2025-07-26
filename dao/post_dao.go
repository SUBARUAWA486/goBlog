package dao

import (
	"goBlog/config"
	"goBlog/model"

	"gorm.io/gorm"
)

func CreatePost(post *model.Post) error {
	return config.DB.Create(post).Error
}

func GetPostByID(id string) (*model.Post, error) { 
	var post model.Post
	err := config.DB.Where("id = ?", id).First(&post).Error
	return &post, err
}

func DeletePost(id string, userID string) error { 
	return config.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Post{}).Error
}

func UpdatePost(post *model.Post) error {
	return config.DB.Save(post).Error
}

func SearchPostsByTitle(title string) ([]model.Post, error) {
	var posts []model.Post
	err := config.DB.Where("title LIKE ?", "%"+title+"%").Find(&posts).Error
	return posts, err
}

func IncrementPostViewCountInDB(postID string, count int) error {
	return config.DB.Model(&model.Post{}).
		Where("id = ?", postID).
		Update("view_count", gorm.Expr("view_count + ?", count)).
		Error
}
