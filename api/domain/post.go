package domain

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID          uint64         `json:"-" gorm:"primaryKey;autoIncrement"`
	Slug        string         `json:"slug" binding:"required" gorm:"type:varchar(128);uniqueIndex"`
	Title       string         `json:"title" binding:"required" gorm:"type:varchar(256)"`
	Content     string         `json:"content"`
	HTML        string         `json:"html" gorm:"-"`
	IsPublished bool           `json:"is_published"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type PostUsecase interface {
	Fetch(ctx context.Context, cursor string, num int) (posts []Post, nextCursor string, err error)
	GetBySlug(ctx context.Context, slug string) (*Post, error)
	Store(ctx context.Context, post *Post) error
	UpdateBySlug(ctx context.Context, slug string, post *Post) error
}

type PostRepository interface {
	Fetch(ctx context.Context, cursor string, num int) ([]Post, error)
	NextCursor(ctx context.Context, slug string) (string, error)
	GetBySlug(ctx context.Context, slug string) (*Post, error)
	Store(ctx context.Context, post *Post) error
	UpdateBySlug(ctx context.Context, slug string, post *Post) error
}
