package http

import (
	"github.com/eindex/qing-zhuo/api/domain"
	"github.com/eindex/qing-zhuo/api/post/delivery/http/middleware"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	PostUsecase domain.PostUsecase
}

func NewPostHandler(rg *gin.RouterGroup, us domain.PostUsecase) {
	handler := &PostHandler{
		PostUsecase: us,
	}
	rg.GET(":slug", handler.GetBySlug)
	rg.GET("", handler.FetchPost)
	rg.POST("", middleware.PermissionCheck(middleware.POST_EDITOR), handler.Store)
	rg.PUT(":slug", middleware.PermissionCheck(middleware.POST_EDITOR), handler.UpdateBySlug)
}

func (p *PostHandler) FetchPost(c *gin.Context) {
	posts, nextCursor, err := p.PostUsecase.Fetch(c, "", 100)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{"data": posts, "next_cursor": nextCursor})
}
func (p *PostHandler) GetBySlug(c *gin.Context) {
	post, err := p.PostUsecase.GetBySlug(c, c.Param("slug"))
	if err != nil {
		return
	}
	c.JSON(200, post)
}
func (p *PostHandler) Store(c *gin.Context) {
	var postRequest domain.Post
	if err := c.BindJSON(&postRequest); err != nil {
		c.JSON(400, err)
		return
	}
	if err := p.PostUsecase.Store(c, &postRequest); err != nil {
		c.JSON(500, "server error")
		return
	}
	c.JSON(200, "success")
}

func (p *PostHandler) UpdateBySlug(c *gin.Context) {
	var postRequest domain.Post
	if err := c.BindJSON(&postRequest); err != nil {
		c.JSON(400, err)
		return
	}
	if err := p.PostUsecase.UpdateBySlug(c, c.Param("slug"), &postRequest); err != nil {
		c.JSON(500, "server error")
		return
	}
	c.JSON(200, "success")
}
