package dao

import (
	"goBlog/config"
	"goBlog/model"
)

func CreateComment(comment *model.Comment) error {
	return config.DB.Create(comment).Error
}

func GetCommentsByPostID(postID string) ([]model.Comment, error) { 
	var comments []model.Comment
	err := config.DB.Where("post_id = ?", postID).Order("created_at desc").Find(&comments).Error
	return comments, err
}