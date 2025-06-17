package controllers

import (
	"blog_system/models/model"
	"blog_system/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(DB *gorm.DB) PostController {
	return PostController{DB}
}

func (pc *PostController) CreatePost(ctx *gin.Context) {
	userID := ctx.MustGet("userID")
	uid, _ := userID.(interface{}).(float64)

	var payload *utils.CreatePostReq
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	newPost := model.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  int32(uid),
	}

	result := pc.DB.Create(&newPost)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"post": newPost}})
}

func (pc *PostController) UpdatePost(ctx *gin.Context) {
	userID := ctx.MustGet("userID")
	uid, _ := userID.(interface{}).(float64)
	postId := ctx.Param("postId")

	var payload *utils.UpdatePostReq
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedPost model.Post
	result := pc.DB.First(&updatedPost, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	if updatedPost.UserID != int32(uid) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not the owner of this post"})
		return
	}

	postToUpdate := model.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  updatedPost.UserID,
	}

	pc.DB.Model(&updatedPost).Updates(postToUpdate)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"post": updatedPost}})
}

func (pc *PostController) FindPostById(ctx *gin.Context) {
	postId := ctx.Param("postId")

	var post model.Post
	result := pc.DB.First(&post, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"post": post}})
}

func (pc *PostController) FindPosts(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var posts []model.Post
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&posts)
	if results.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": results.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(posts), "data": gin.H{"posts": posts}})
}

func (pc *PostController) DeletePost(ctx *gin.Context) {
	postId := ctx.Param("postId")
	userID := ctx.MustGet("userID")
	uid, _ := userID.(interface{}).(float64)

	var post model.Post
	result := pc.DB.First(&post, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	if post.UserID != int32(uid) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not the owner of this post"})
		return
	}

	result = pc.DB.Delete(&model.Post{}, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
