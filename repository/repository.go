package repository

import "database/sql"

type DatabaseRepo interface {}

type repo struct {
	DB *sql.DB
}

func New(conn *sql.DB) DatabaseRepo {
	return &repo{
		DB: conn,
	}
}