package domain

import "context"

type Author struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AuthorRepository interface {
	GetById(ctx context.Context, id int64) (Author, error)
}
