package http

import (
	"github.com/eindex/qing-zhuo/api/domain"
	"github.com/gin-gonic/gin"
)

type ArchiveHandler struct {
	PostUsecase domain.PostUsecase
}

func NewArchiveHandler(rg *gin.RouterGroup, us domain.PostUsecase) {
	handler := &ArchiveHandler{
		PostUsecase: us,
	}
	rg.GET("", handler.FetchArchive)
}

func (p *ArchiveHandler) FetchArchive(c *gin.Context) {
	posts, nextCursor, err := p.PostUsecase.Fetch(c, "", 100)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{"data": posts, "next_cursor": nextCursor})
}
