package service

import (
	"errors"

	"goBlog/model"
	"goBlog/dao"
	"goBlog/utils"
)

func CreateComment(postID, userID, content string) (error) {

	if len(content)>100 || len(content)<1 { 
		return errors.New("评论内容长度不符")
	}

	user, err := dao.GetUserByID(userID)
	if err != nil { 
		return errors.New("用户不存在")
	}

	_, err =dao.GetPostByID(postID)
	if err != nil {
		return errors.New("文章不存在")
	}

	comment := &model.Comment{
		ID: utils.GenerateCommentID(),
		PostID: postID,
		UserID: userID,
		Nickname: user.Nickname,
		Content: content,
	}

	err = dao.CreateComment(comment)
	if err != nil { 
		return errors.New("创建评论失败")
	}
	return nil
}

func GetComments(postID string) ([]model.Comment, error) {
	return dao.GetCommentsByPostID(postID)
}