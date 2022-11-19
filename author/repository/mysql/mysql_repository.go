package mysql

import (
	"context"
	"database/sql"

	"github.com/grimoh/go-test-server/domain"
)

type mySQLAuthorRepo struct {
	DB *sql.DB
}

func NewMySQLAuthorRepository(db *sql.DB) domain.AuthorRepository {
	return &mySQLAuthorRepo{
		DB: db,
	}
}

func (m *mySQLAuthorRepo) getOne(ctx context.Context, query string, args ...interface{}) (res domain.Author, err error) {
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return domain.Author{}, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	res = domain.Author{}

	err = row.Scan(
		&res.Id,
		&res.Name,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	return
}

func (m *mySQLAuthorRepo) GetById(ctx context.Context, id int64) (domain.Author, error) {
	query := `SELECT id, name, created_at, updated_at FROM author WHERE id=?`
	return m.getOne(ctx, query, id)
}
