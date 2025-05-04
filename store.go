package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}
type Article struct {
	ID          string
	Title       string
	Description string
	URL         string
	CreatedAt   string
}

func (s *Store) Init() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://admin:adminpassword@localhost/readercliDB?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil

}
func (s *Store) Create(ctx context.Context, article *Article) error {
	query := `
	INSERT INTO articles(title,description,url)
	VALUES($1,$2,$3)
	RETURNING id, created_at
	`
	err := s.db.QueryRowContext(ctx, query,
		article.Title,
		article.Description,
		article.URL).Scan(&article.ID, article.CreatedAt)
	if err != nil {
		return err
	}
	return nil

}
func (s *Store) GetArticles(ctx context.Context) ([]Article, error) {
	query := `
	SELECT * FROM articles
	`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	articles := []Article{}
	for rows.Next() {
		var a Article
		err := rows.Scan(&a.ID, &a.CreatedAt, &a.Title, &a.Description, &a.URL)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}
	return articles, nil

}
