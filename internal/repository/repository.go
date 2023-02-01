package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Transaction interface {
}

type Repository struct {
	Authorization
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
