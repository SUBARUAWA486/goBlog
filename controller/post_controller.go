package controller

import (
	"goBlog/response"
	"goBlog/service"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var req struct {
		Title   string `json:"title" binding:"required,max=20"`
		Content string `json:"content" binding:"required,max=300"`
		Cover   string `json:"cover" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.CodeInvalidParam, "参数错误: "+err.Error())
		return
	}

	userID := c.GetString("user_id")
	if userID == "" { 
		response.Error(c, response.CodeAuthFailed, "用户未登录")
		return
	}

	post, err := service.CreatePost(req.Title, req.Content, req.Cover, userID)
	if err != nil {
		response.Error(c, response.CodeServerError, "创建帖子失败："+err.Error())
		return
	}

	response.Success(c, gin.H{"post_id": post.ID})
}

func SearchPosts(c *gin.Context) {

	title := c.Query("title")
	if title == "" {
		response.Error(c, response.CodeNotFound, "title 缺失")
		return
	}

	posts, err := service.SearchPosts(title)
	if err != nil {
		response.Error(c, response.CodeServerError, "查询文章失败："+err.Error())
		return
	}
	response.Success(c, posts)
}

func GetPostDetails(c *gin.Context) {
	
	postID := c.Param("post_id")
	if postID == "" {
		response.Error(c, response.CodeAuthFailed, "未登录")
		return
	}

	userID := c.GetString("user_id")
	if userID == "" {
		response.Error(c, response.CodeNotFound, "缺少 user_id")
		return
	}

	postDetail, err := service.GetPostDetailWithExtras(postID, userID)
	if err != nil {
		response.Error(c, response.CodeServerError, "获取失败："+err.Error())
		return
	}

	//异步记录浏览量
	go service.AsyncAddPostView(postID)

	response.Success(c, postDetail)
}
