package main

import "database/sql"

type Store struct {
	db *sql.DB
}
type Article struct{
	ID string
	Title string
	Description string
	URL string
	CreatedAt string

}

func (s *Store) Init() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://admin:adminpassword@localhost/readercliDB?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil

}
