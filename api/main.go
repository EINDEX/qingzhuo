package main

import (
	"fmt"
	"github.com/eindex/qing-zhuo/api/domain"
	"github.com/eindex/qing-zhuo/api/post/delivery/http"
	"github.com/eindex/qing-zhuo/api/post/delivery/http/middleware"
	mysql2 "github.com/eindex/qing-zhuo/api/post/repository/mysql"
	"github.com/eindex/qing-zhuo/api/post/usercase"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

func main() {
	dsn := os.Getenv("DSN")
	DBConn, _ := gorm.Open(mysql.Open(dsn))
	err := DBConn.AutoMigrate(&domain.Post{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	router := gin.Default()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	pr := mysql2.NewMysqlPostRepository(DBConn)

	api := router.Group("api", middleware.PermissionApplyMiddleware())

	pu := usercase.NewPostUsecase(pr, timeoutContext)
	http.NewPostHandler(api.Group("posts"), pu)

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

	router.Run()
}
