package usercase

import (
	"bytes"
	"context"
	"github.com/eindex/qing-zhuo/api/domain"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/util"
	"time"
)

const CUT_SIZE int = 140

var md = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		extension.Footnote,
		extension.TaskList,
	),
)

type postUsecase struct {
	postRepo       domain.PostRepository
	contextTimeout time.Duration
}

func NewPostUsecase(p domain.PostRepository, timeout time.Duration) domain.PostUsecase {
	return &postUsecase{
		postRepo:       p,
		contextTimeout: timeout,
	}
}

func (p *postUsecase) Fetch(ctx context.Context, cursor string, num int) (posts []domain.Post, nextCursor string, err error) {
	posts, err = p.postRepo.Fetch(ctx, cursor, num)
	if err != nil {
		return
	}
	for i := range posts {
		posts[i].HTML = MarkdownRender(Summary(posts[i].Content))
	}
	return
}

func (p *postUsecase) GetBySlug(ctx context.Context, slug string) (*domain.Post, error) {
	post, err := p.postRepo.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	post.HTML = MarkdownRender(post.Content)
	return post, nil
}

func (p *postUsecase) Store(ctx context.Context, post *domain.Post) error {
	err := p.postRepo.Store(ctx, post)
	return err
}

func (p *postUsecase) UpdateBySlug(ctx context.Context, slug string, post *domain.Post) error {
	return p.postRepo.UpdateBySlug(ctx, slug, post)
}

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
