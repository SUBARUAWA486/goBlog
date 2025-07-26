package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) { 
	c.JSON(http.StatusOK, Result{
		Code: 0,
		Msg:  "操作成功",
		Data: data,
	})
}

func Error(c *gin.Context, code int, msg string) { 
	c.JSON(http.StatusOK, Result{
		Code: code,
		Msg:  msg,
	})
}