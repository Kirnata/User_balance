package repository

import (
	"github.com/Kirnata/User_balance"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user User_balance.User) (int, error)
	GetUser(username, password string) (User_balance.User, error)
}

type Payments interface {
}

type Balance interface {
	GetBalance(id int) (float32, error)
}

type Repository struct {
	Authorization
	Payments
	Balance
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Balance:       NewBalanceServicePostgres(db),
	}
}
