package utils

import (
	"errors"
	"time"
	"goBlog/config"

	"github.com/golang-jwt/jwt"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID   string `json:"user_id"`
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

func GenerateToken(userID, nickname string) (string, error) {

	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := Claims{
		UserID:   userID,
		Nickname: nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "goblog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return "", err
	}

	// 存入 Redis（以 userID 为 key）
	key := "login:token:" + userID
	err = config.RedisClient.Set(config.Ctx, key, tokenStr, 7*24*time.Hour).Err()
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ParseToken(tokenStr string) (*Claims, error) {
	
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecret), nil
	})
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// 进一步从 Redis 中校验 token 是否有效
		key := "login:token:" + claims.UserID
		storedToken, err := config.RedisClient.Get(config.Ctx, key).Result()
		if err != nil || storedToken != tokenStr {
			return nil, errors.New("token失效或已被踢出")
		}
		return claims, nil
	}
	return nil, err
}

func ExtractUserIDFromToken(c *gin.Context) (string, error) {
	uid, exists := c.Get("user_id")
	if !exists {
		return "", errors.New("用户 ID 不存在于上下文中")
	}
	userID, ok := uid.(string)
	if !ok {
		return "", errors.New("用户 ID 类型断言失败")
	}
	return userID, nil
}