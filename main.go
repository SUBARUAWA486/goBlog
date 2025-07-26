package main

import (
	"goBlog/config"
	"goBlog/model"
	"goBlog/router"
	"goBlog/service"
	"time"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载 .env 文件（数据库配置、JWT 密钥等）
	config.LoadEnv()

	// 初始化数据库连接
	config.InitDB()
	config.InitRedis()

	// 自动迁移数据表（如果不存在则创建）
	err := config.DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{},&model.Stars{})
	if err != nil {
		log.Fatal("数据库迁移失败：", err)
	}

	// 初始化日志器
	config.InitLogger()

	go func() {
	for {
		time.Sleep(5 * time.Minute)
		service.SyncViewCountsToDB()
	}
	}()

	// 启动路由
	r := router.SetupRouter()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"msg":    "页面不存在",
			"path":   c.Request.URL.Path,
			"method": c.Request.Method,
		})
	})

	// 启动服务，监听 8080 端口
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("服务启动失败：", err)
	}
}
