package controller

import (
	"goBlog/response"
	"goBlog/service"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	var req struct {
		Content string `json:"content" binding:"required,max=100"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.CodeInvalidParam, "参数错误："+err.Error())
		return
	}

	postID := c.Param("post_id")
    if postID == "" {
        response.Error(c, response.CodeNotFound, "post_id 缺失")
        return
    }

	userID := c.GetString("user_id")
	if userID == "" {
		response.Error(c, response.CodeAuthFailed, "用户未登录")
		return
	}
	
	if err := service.CreateComment(postID, userID, req.Content); err != nil {
		response.Error(c, response.CodeServerError, "创建评论失败："+err.Error())
		return
	}

	response.Success(c, "创建评论成功")
}

func GetComments(c *gin.Context) {
	postID := c.Param("post_id")
	if postID == "" {
		response.Error(c, response.CodeNotFound, "缺少 post_id")
		return
	}

	comments, err := service.GetComments(postID)
	if err != nil {
		response.Error(c, response.CodeServerError, "获取评论失败："+err.Error())
		return
	}

	response.Success(c, comments)
}
