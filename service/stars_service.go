package service

import (
	"errors"

	"goBlog/dao"
	"goBlog/model"
)

func AddPostStar(userID, postID string) error {
	// 避免重复收藏
	exists, err := dao.IsStarred(userID, postID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("您已收藏过该文章")
	}

	err = dao.AddStars(userID, postID)
	if err != nil {
		return err
	}

	dao.IncrementStars(postID)
	return nil
}

func RemovePostStar(userID, postID string) error {
	exists, err := dao.IsStarred(userID, postID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("您尚未收藏该文章")
	}

	err = dao.RemoveStar(userID, postID)
	if err != nil {
		return err
	}

	dao.DecrementStars(postID)
	return nil
}


func GetStarredPosts(userID string) ([]model.Post, error) {
	
	posts, err := dao.GetUserStars(userID)
	if err != nil {
		return nil, err
	}
	return posts, nil
}