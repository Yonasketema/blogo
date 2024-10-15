package models

import (
	"database/sql"
	"errors"
	"time"
)

type Blog struct {
	Id      int
	Title   string
	Content string
	Created time.Time
}

type BlogModel struct {
	DB *sql.DB
}

func (m *BlogModel) GetOneBlog(id int) (Blog, error) {
	sqls := `SELECT id, title, content, created FROM blogs
   WHERE id = ?`
	row := m.DB.QueryRow(sqls, id)

	var b Blog

	err := row.Scan(&b.Id, &b.Title, &b.Content, &b.Created)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return Blog{}, ErrNoRecord
		} else {
			return Blog{}, err
		}

	}

	return b, nil

}

func (m *BlogModel) GetAllBlog() ([]Blog, error) {

	sqls := `SELECT id, title, content, created  FROM blogs ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(sqls)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []Blog

	for rows.Next() {
		var b Blog
		err := rows.Scan(&b.Id, &b.Title, &b.Content, &b.Created)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (m *BlogModel) InsertBlog(title string, content string) (int, error) {

	sqls := `INSERT INTO blogs (title, content, created)
			 VALUES(?, ?, UTC_TIMESTAMP())`

	result, err := m.DB.Exec(sqls, title, content)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}
