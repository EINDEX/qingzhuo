package mysql

import (
	"context"
	"github.com/eindex/qing-zhuo/api/domain"
	"gorm.io/gorm"
)

type mysqlPostRepository struct {
	Conn *gorm.DB
}

func NewMysqlPostRepository(Conn *gorm.DB) domain.PostRepository {
	return &mysqlPostRepository{Conn}
}

func (m mysqlPostRepository) Fetch(ctx context.Context, cursor string, num int) (posts []domain.Post, err error) {
	if err := m.Conn.Order("created_at desc").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (m mysqlPostRepository) NextCursor(ctx context.Context, slug string) (string, error) {
	// todo
	return "", nil
}

func (m *mysqlPostRepository) GetBySlug(ctx context.Context, slug string) (post *domain.Post, err error) {
	m.Conn.First(&post, "slug = ?", slug)
	return
}

func (m mysqlPostRepository) Store(ctx context.Context, post *domain.Post) (err error) {
	if err = m.Conn.Create(post).Error; err != nil {
		return
	}
	return
}

func (m mysqlPostRepository) UpdateBySlug(ctx context.Context, slug string, post *domain.Post) (err error) {
	if err = m.Conn.Model(&domain.Post{}).Where("slug = ?", slug).Omit("id").Updates(post).Error; err != nil {
		return
	}
	return
}
