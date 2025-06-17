package task3

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 进阶gorm
/*
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/

type User struct {
	gorm.Model
	Name      string
	PostCount int
	Posts     []Post
}

type Post struct {
	gorm.Model
	Title         string
	Content       string
	UserID        uint
	CommentCount  int
	CommentStatus string
	Comments      []Comment
}

type Comment struct {
	gorm.Model
	Content string
	PostID  uint
}

func gorm1() {
	// 创建数据库连接，这里使用 SQLite 作为示例
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移，创建表
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic("failed to migrate database")
	}

	// 进行其他操作
}

/*
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/

func getAllCommentsByUid(db *gorm.DB, userID uint32) {
	// 查询某个用户发布的所有文章及其对应的评论信息
	var user User
	err := db.Preload("Posts.Comments").First(&user, userID).Error
	if err != nil {
		fmt.Println("查询错误:", err)
		return
	}

	fmt.Printf("用户: %s\n", user.Name)
	for _, post := range user.Posts {
		fmt.Printf("文章: %s\n", post.Title)
		for _, comment := range post.Comments {
			fmt.Printf("评论: %s\n", comment.Content)
		}
	}
}

func findMostCommentedPost(db *gorm.DB) {
	type Result struct {
		PostID       uint
		CommentCount int64
	}

	var result Result
	err := db.Table("comments").
		Select("post_id, count(*) as comment_count").
		Group("post_id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&result).Error
	if err != nil {
		fmt.Println("查询错误:", err)
		return
	}

	var mostCommentedPost Post
	err = db.First(&mostCommentedPost, result.PostID).Error
	if err != nil {
		fmt.Println("不能找到有最多评论的文章:", err)
		return
	}
	fmt.Printf("评论最多的文章: %s (评论数: %d)\n", mostCommentedPost.Title, result.CommentCount)
}

/*
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为0，则更新文章的评论状态为 "无评论"。
*/

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 更新用户的 PostCount 字段
	err = tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
	return
}

func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	var count int64

	// 查找当前文章剩余的评论数量
	err = tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error
	if err != nil {
		return err
	}

	// 如果评论数量为 1（因为当前评论还未真正删除），这意味着删除后数量将变为 0
	if count == 1 {
		// 更新文章的 CommentStatus 字段为 "无评论"
		err = tx.Model(&Post{}).Where("id = ?", c.PostID).UpdateColumn("comment_status", "无评论").Error
		if err != nil {
			return err
		}
	}

	return nil
}
