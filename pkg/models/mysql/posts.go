package mysql

import (
	"database/sql"
	"errors"
	"github.com/Mike-95/blog_web/pkg/models"
)

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) Insert(title, content, category string) (int, error) {
	stmt := `insert into posts (title, content, created, category) values (?, ?, UTC_TIMESTAMP(),?)`

	result, err := m.DB.Exec(stmt, title, content, category)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *PostModel) Get(id int) (*models.Post, error) {
	stmt := `select id, title, content, created, category from posts where id = ?`
	row := m.DB.QueryRow(stmt, id)

	s := &models.Post{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Category)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *PostModel) Latest() ([]*models.Post, error) {
	stmt := `select id, title, content, created, category from posts order by created desc LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	posts := []*models.Post{}
	for rows.Next() {
		s := &models.Post{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Category)
		if err != nil {
			return nil, err
		}
		posts = append(posts, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}
