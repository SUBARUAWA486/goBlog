package router

import (
	"time"
	"goBlog/controller"
	"goBlog/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func SetupRouter() *gin.Engine {
	
	r := gin.Default()

	// CORS 配置
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")

	// 不需要认证的接口
	api.POST("/register", controller.Register)
	api.POST("/login", controller.Login)

	// 需要 JWT 认证的接口
	auth := api.Group("/")
	auth.Use(middleware.JWTAuth())
	{
		
		auth.POST("/logout", controller.Logout)


		auth.POST("/posts", controller.CreatePost)
		auth.GET("/posts", controller.SearchPosts)
		auth.GET("/posts/:post_id", controller.GetPostDetails)


		auth.POST("/posts/:post_id/comments", controller.AddComment)
		auth.GET("/posts/:post_id/comments", controller.GetComments)

		
		auth.POST("/posts/:post_id/star", controller.StarPost)
		auth.DELETE("/posts/:post_id/star", controller.UnstarPost)
		auth.GET("/star", controller.GetStarredPosts)

	}
	return r
}