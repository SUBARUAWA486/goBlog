package controller

import (
	"goBlog/response"
	"goBlog/service"
	"goBlog/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req struct {
		Account  string `json:"account" binding:"required,max=9"`
		Nickname string `json:"nickname" binding:"required,max=12"`
		Password string `json:"password" binding:"required,min=6,max=12"`
		Avatar   string `json:"avatar" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.CodeInvalidParam, "参数错误:"+err.Error())
		return
	}

	newID, err := service.RegisterUser(req.Account, req.Nickname, req.Password, req.Avatar)
	if err != nil {
		response.Error(c, response.CodeServerError, "注册失败:"+err.Error())
		return
	}

	response.Success(c, gin.H{"user_id": newID})
}

func Login(c *gin.Context) {
	var req struct {
		Account string `json:"account"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.CodeInvalidParam, "参数错误")
		return
	}

	user, err := service.LoginUser(req.Account, req.Password)
	if err != nil {
		response.Error(c, response.CodeAuthFailed, "登录失败: "+err.Error())
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Nickname)
	if err != nil {
		response.Error(c, response.CodeServerError, "生成 token 失败: "+err.Error())
		return
	}
	response.Success(c, gin.H{
		"message": "登录成功",
		"token":   token,
		"account": user.Account,
	})
}

func Logout(c *gin.Context) {
	userID, err := utils.ExtractUserIDFromToken(c)
	if err != nil {
		response.Error(c, response.CodeAuthFailed, "无效的 token")
		return
	}

	err = service.LogoutUser(userID)
	if err != nil {
		response.Error(c, response.CodeServerError, "退出登录失败: "+err.Error())
		return
	}

	response.Success(c, "退出登录成功")
}