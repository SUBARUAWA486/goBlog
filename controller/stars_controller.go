package controller

import (
	"goBlog/response"
	"goBlog/service"
	
	"github.com/gin-gonic/gin"
)

func StarPost(c *gin.Context) {

	userID := c.GetString("user_id")
	if userID == "" {
		response.Error(c, response.CodeAuthFailed, "未登录")
		return
	}

	postID := c.Param("post_id")
	if postID == "" {
		response.Error(c, response.CodeNotFound, "post_id 缺失")
		return
	}

	err := service.AddPostStar(userID, postID)
	if err != nil {
		response.Error(c, response.CodeServerError, err.Error())
		return
	}

	response.Success(c, "收藏成功")
}

func UnstarPost(c *gin.Context) {

	userID := c.GetString("user_id")
	if userID == "" {
		response.Error(c, response.CodeAuthFailed, "未登录")
		return
	}

	postID := c.Param("post_id")
	if postID == "" {
		response.Error(c, response.CodeNotFound, "post_id 缺失")
		return
	}

	err := service.RemovePostStar(userID, postID)
	if err != nil {
		response.Error(c, response.CodeServerError, err.Error())
		return
	}

	response.Success(c, "取消收藏成功")
}

func GetStarredPosts(c *gin.Context) {

	userID := c.GetString("user_id")
	if userID == "" {
		response.Error(c, response.CodeAuthFailed, "未登录")
		return
	}

	posts, err := service.GetStarredPosts(userID)
	if err != nil {
		response.Error(c, response.CodeServerError, err.Error())
		return
	}

	response.Success(c, posts)
}