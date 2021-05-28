package main

import (
	"bytes"
	"fmt"
	"github.com/eindex/qing-zhuo/api/delivery/http/middleware"
	"github.com/eindex/qing-zhuo/api/domain"
	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

const CUT_SIZE int = 140

func (p *domain.CreateUpdatePostRequest) getPost() *Post {
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
		content = string([]rune(content)[:CUT_SIZE]) + "\n ..."
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

	api := router.Group("api", middleware.PermissionApplyMiddleware())
	{
		posts := api.Group("posts")
		{

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
