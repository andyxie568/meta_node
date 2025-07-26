package main

import (
	"blog_system/controllers"
	middleware "blog_system/middlewares"
	"blog_system/models"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func init() {
	// 加载环境变量
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "123456")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "blog")
}

func main() {
	// 初始化数据库
	models.ConnectDB()

	// 创建Gin路由引擎
	router := gin.Default()

	// 注册中间件
	router.Use(middleware.LoggerMiddleware())

	// 初始化控制器
	authController := controllers.NewAuthController(models.DB)
	postController := controllers.NewPostController(models.DB)
	commentController := controllers.NewCommentController(models.DB)

	// 认证路由
	router.POST("/api/auth/register", authController.SignUpUser)
	router.POST("/api/auth/login", authController.SignInUser)
	router.POST("/api/auth/logout", middleware.AuthMiddleware(), authController.LogoutUser)

	// 文章路由
	router.POST("/api/posts", middleware.AuthMiddleware(), postController.CreatePost)
	router.PUT("/api/posts/:postId", middleware.AuthMiddleware(), postController.UpdatePost)
	router.GET("/api/posts/:postId", postController.FindPostById)
	router.GET("/api/posts", postController.FindPosts)
	router.DELETE("/api/posts/:postId", middleware.AuthMiddleware(), postController.DeletePost)

	// 评论路由
	router.POST("/api/posts/:postId/comments", middleware.AuthMiddleware(), commentController.CreateComment)
	router.GET("/api/posts/:postId/comments", commentController.GetComments)
	router.DELETE("/api/comments/:commentId", middleware.AuthMiddleware(), commentController.DeleteComment)

	// 启动服务器
	log.Fatal(router.Run(":9526"))
}
