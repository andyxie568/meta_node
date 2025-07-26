package controllers

import (
	"blog_system/models/model"
	"blog_system/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct {
	DB *gorm.DB
}

func NewCommentController(DB *gorm.DB) CommentController {
	return CommentController{DB}
}

func (cc *CommentController) CreateComment(ctx *gin.Context) {
	userID := ctx.MustGet("userID")
	uid, _ := userID.(interface{}).(float64)

	postId := ctx.Param("postId")

	var payload *utils.CreateCommentReq
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var post model.Post
	result := cc.DB.First(&post, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that id exists"})
		return
	}

	newComment := model.Comment{
		Content: payload.Content,
		UserID:  int32(uid),
		PostID:  int32(post.ID),
	}

	result = cc.DB.Create(&newComment)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"comment": newComment}})
}

func (cc *CommentController) GetComments(ctx *gin.Context) {
	postId := ctx.Param("postId")

	var comments []model.Comment
	result := cc.DB.Where("post_id = ?", postId).Find(&comments)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(comments), "data": gin.H{"comments": comments}})
}

func (cc *CommentController) DeleteComment(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	userID := ctx.MustGet("userID")
	uid, _ := userID.(interface{}).(float64)

	var comment model.Comment
	result := cc.DB.First(&comment, "id = ?", commentId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No comment with that id exists"})
		return
	}

	if comment.UserID != int32(uid) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not the owner of this comment"})
		return
	}

	result = cc.DB.Delete(&model.Comment{}, "id = ?", commentId)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
