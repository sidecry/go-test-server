package domain

import (
	"context"
	"time"
)

type Article struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    Author    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ArticleUseCase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Article, string, error)
	GetById(ctx context.Context, id int64) (Article, error)
	Update(ctx context.Context, ar *Article) error
	GetByTitle(ctx context.Context, title string) (Article, error)
	Store(context.Context, *Article) error
	Delete(ctx context.Context, id int64) error
}

type ArticleRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Article, nextCursor string, err error)
	GetById(ctx context.Context, id int64) (Article, error)
	GetByTitle(ctx context.Context, title string) (Article, error)
	Update(ctx context.Context, ar *Article) error
	Store(ctx context.Context, a *Article) error
	Delete(ctx context.Context, id int64) error
}
