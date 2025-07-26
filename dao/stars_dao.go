package dao

import (
	"goBlog/config"
	"goBlog/model"

	"gorm.io/gorm"
)

func AddStars(userID, postID string) error {

	stars := model.Stars{
		ID:     userID + "_" + postID,
		UserID: userID,
		PostID: postID,
	}
	return config.DB.Create(&stars).Error
}

func IsStarred(userID, postID string) (bool, error) {
	var count int64
	err := config.DB.Model(&model.Stars{}).
		Where("user_id = ? AND post_id = ?", userID, postID).
		Count(&count).Error
	return count > 0, err
}

func RemoveStar(userID, postID string) error {
	return config.DB.Delete(&model.Stars{}, "id = ?", userID+"_"+postID).Error
}

func GetUserStars(userID string) ([]model.Post, error) {
	var posts []model.Post
	err := config.DB.
		Joins("JOIN stars ON stars.post_id = posts.id").
		Where("stars.user_id = ?", userID).
		Find(&posts).Error
	return posts, err
}

func IncrementStars(postID string) {
	config.DB.Model(&model.Post{}).
		Where("id = ?", postID).
		Update("stars", gorm.Expr("stars + 1"))
}

func DecrementStars(postID string) {
	config.DB.Model(&model.Post{}).
		Where("id = ?", postID).
		Update("stars", gorm.Expr("CASE WHEN stars > 0 THEN stars - 1 ELSE 0 END"))
}