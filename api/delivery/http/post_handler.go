package http

import (
	"github.com/eindex/qing-zhuo/api/delivery/http/middleware"
	"github.com/eindex/qing-zhuo/api/domain"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	PostUseCase domain.PostUseCase
}

func NewPostHandler(rg *gin.RouterGroup, us domain.PostUseCase) {
	handler := &PostHandler{
		PostUseCase: us,
	}
	rg.GET(":slug", handler.GetBySlug)
	rg.GET("", handler.FetchPost)
	rg.POST("", middleware.PermissionCheck(middleware.POST_EDITOR), handler.Store)
	rg.PUT(":slug", middleware.PermissionCheck(middleware.POST_EDITOR), handler.UpdateBySlug)
}

func (p *PostHandler) FetchPost(c *gin.Context) {
	var posts []Post
	DB.Order("created_at desc").Find(&posts)
	for i := range posts {
		posts[i].HTML = MarkdownRender(Summary(posts[i].Content))
	}
	c.JSON(200, posts)
}
func (p *PostHandler) GetBySlug(c *gin.Context) {
	var post Post
	DB.First(&post, "slug = ?", c.Param("slug"))
	post.HTML = MarkdownRender(post.Content)
	c.JSON(200, post)
}
func (p *PostHandler) Store(c *gin.Context) {
	var postRequest domain.CreateUpdatePostRequest
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
}

func (p *PostHandler) UpdateBySlug(c *gin.Context) {
	var postRequest domain.CreateUpdatePostRequest
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
}
