package middleware

import (
	"strings"
	"goBlog/response"
	"goBlog/utils"
	
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			response.Error(c, response.CodeAuthFailed, "权限不足,未携带token")
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			response.Error(c, response.CodeAuthFailed, "token无效或已过期")
			c.Abort()
			return
		}

		// 保存用户信息到上下文
		c.Set("user_id", claims.UserID)
		c.Set("nickname", claims.Nickname)

		c.Next()
	}
}
