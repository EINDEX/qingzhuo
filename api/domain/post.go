package domain

import "context"

type Post struct {
	Slug    string `json:"slug" binding:"required"  validator:""`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
	Publish bool   `json:"publish"`
}

type PostUseCase interface {
	Fetch(ctx context.Context, cursor string, slug string) ([]Post, string, error)
}

type PostRepository interface {
}
