package main

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/eindex/qing-zhuo/api/premissions"
	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const CUT_SIZE int = 270

type Post struct {
	ID          uint64         `json:"id" gorm:"primaryKey;autoIncrement"`
	Slug        string         `json:"slug" gorm:"type:varchar(128);uniqueIndex"`
	Title       string         `json:"title" gorm:"type:varchar(256)"`
	Content     string         `json:"content"`
	HTML        string         `gorm:"-" json:"html"`
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

var md = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		extension.Footnote,
		extension.TaskList,
	),
)

func MarkdownRender(markdown string) (html string) {

	var buf bytes.Buffer
	if err := md.Convert(util.StringToReadOnlyBytes(markdown), &buf); err != nil {
		return
	}
	html = buf.String()
	return
}

func Summary(content string) string {
	if len(content) > CUT_SIZE {
		content = content[:CUT_SIZE] + "\n ..."
	}
	return content
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
				post.HTML = MarkdownRender(post.Content)
				c.JSON(200, post)
			})
			posts.GET("", func(c *gin.Context) {
				var posts []Post
				DB.Order("created_at desc").Find(&posts)
				for i := range posts {
					posts[i].HTML = MarkdownRender(Summary(posts[i].Content))
				}
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
		archives := api.Group("archives")
		{
			archives.GET("", func(c *gin.Context) {
				var posts []Post
				DB.Select("created_at, slug, title").Order("created_at desc").Find(&posts)
				for i := range posts {
					posts[i].HTML = MarkdownRender(posts[i].Content)
				}
				postByYear := make(map[int][]Post)
				for _, post := range posts {
					year := post.CreatedAt.Year()
					postsInYear, ok := postByYear[year]
					if !ok {
						postsInYear = make([]Post, 0)
					}
					postsInYear = append(postsInYear, post)
					postByYear[year] = postsInYear
				}
				c.JSON(200, postByYear)
			})
		}
	}

	router.Run()
}
