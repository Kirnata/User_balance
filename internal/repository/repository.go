package repository

import (
	"github.com/Kirnata/User_balance"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user User_balance.User) (int, error)
	GetUser(username, password string) (User_balance.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
