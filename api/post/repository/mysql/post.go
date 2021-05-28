package mysql

import (
	"gorm.io/gorm"
	"time"
)

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

func (p *CreateUpdatePostRequest) getPost() *Post {
	return &Post{
		Slug:        p.Slug,
		Title:       p.Title,
		Content:     p.Content,
		IsPublished: p.Publish,
	}
}
