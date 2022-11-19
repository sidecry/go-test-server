package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/grimoh/go-test-server/article/repository"
	"github.com/grimoh/go-test-server/domain"
	"go.uber.org/zap"
)

type mySQLArticleRepository struct {
	Conn *sql.DB
}

func NewMySQLArticleRepository(Conn *sql.DB) domain.ArticleRepository {
	return &mySQLArticleRepository{Conn}
}

func (m *mySQLArticleRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Article, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	defer func() {
		if errRow := rows.Close(); err != nil {
			zap.S().Error(errRow)
		}
	}()

	result = make([]domain.Article, 0)
	for rows.Next() {
		t := domain.Article{}
		authorId := int64(0)
		err = rows.Scan(
			&t.Id,
			&t.Title,
			&t.Content,
			&authorId,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			zap.S().Error(err)
			return nil, err
		}

		t.Author = domain.Author{
			Id: authorId,
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mySQLArticleRepository) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Article, nextCursor string, err error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM article WHERE created_at > ? ORDER BY created_at LIMIT ? `

	decodedCursor, err := repository.DecodeCursor(cursor)
	if err != nil && cursor != "" {
		return nil, "", domain.ErrBadParamInput
	}

	res, err = m.fetch(ctx, query, decodedCursor, num)
	if err != nil {
		return nil, "", err
	}

	if len(res) == int(num) {
		nextCursor = repository.EncodeCursor(res[len(res)-1].CreatedAt)
	}

	return
}
func (m *mySQLArticleRepository) GetById(ctx context.Context, id int64) (res domain.Article, err error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM article WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Article{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mySQLArticleRepository) GetByTitle(ctx context.Context, title string) (res domain.Article, err error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM article WHERE title = ?`

	list, err := m.fetch(ctx, query, title)
	if err != nil {
		return
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}
	return
}

func (m *mySQLArticleRepository) Store(ctx context.Context, a *domain.Article) (err error) {
	query := `INSERT  article SET title=? , content=? , author_id=?, updated_at=? , created_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, a.Title, a.Content, a.Author.Id, a.UpdatedAt, a.CreatedAt)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	a.Id = lastID
	return
}

func (m *mySQLArticleRepository) Delete(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM article WHERE id = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", rowsAfected)
		return
	}

	return
}
func (m *mySQLArticleRepository) Update(ctx context.Context, ar *domain.Article) (err error) {
	query := `UPDATE article set title=?, content=?, author_id=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, ar.Title, ar.Content, ar.Author.Id, ar.UpdatedAt, ar.Id)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}
