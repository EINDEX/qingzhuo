package main

import (
	"fmt"
	"os"
	"time"

	"github.com/eindex/qing-zhuo/api/premissions"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Post struct {
	ID          uint64         `json:"id" gorm:"primaryKey;autoIncrement"`
	Slug        string         `json:"slug" gorm:"type:varchar(128);uniqueIndex"`
	Title       string         `json:"title" gorm:"type:varchar(256)"`
	Content     string         `json:"content"`
	IsPublished bool           `json:"is_published"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateUpdatePostRequest struct {
	Slug    string `json:"slug" binding:"required"  validator:""`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
	Publish bool   `json:"publish"`
}

func (p *CreateUpdatePostRequest) getPost() *Post {
	return &Post{
		Slug:        p.Slug,
		Title:       p.Title,
		Content:     p.Content,
		IsPublished: p.Publish,
	}
}

func main() {
	dsn := os.Getenv("DSN")
	DB, _ := gorm.Open(mysql.Open(dsn))
	err := DB.AutoMigrate(&Post{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	router := gin.Default()

	api := router.Group("api", premissions.PremissionApplyMiddleware())
	{
		posts := api.Group("posts")
		{
			posts.GET(":slug", func(c *gin.Context) {
				var post Post
				DB.First(&post, "slug = ?", c.Param("slug"))
				c.JSON(200, post)

			})
			posts.GET("", func(c *gin.Context) {
				var posts []Post
				DB.Find(&posts)
				c.JSON(200, posts)
			})
			posts.POST("", premissions.PremissionCheck(premissions.POST_EDITOR), func(c *gin.Context) {
				var postRequest CreateUpdatePostRequest
				if err := c.BindJSON(&postRequest); err != nil {
					c.JSON(400, err)
					return
				}
				post := postRequest.getPost()
				if err := DB.Create(&post).Error; err != nil {
					c.JSON(400, err)
					return
				}
				c.JSON(200, "success")
			})
			posts.PUT(":slug", premissions.PremissionCheck(premissions.POST_EDITOR), func(c *gin.Context) {
				var postRequest CreateUpdatePostRequest
				if err := c.BindJSON(&postRequest); err != nil {
					c.JSON(400, err)
					return
				}
				post := postRequest.getPost()
				if err := DB.Save(&post).Error; err != nil {
					c.JSON(400, err)
					return
				}
				c.JSON(200, "success")
			})
		}
	}

	router.Run()
}
