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
	CreateSinglePay(id int, input User_balance.SinglePay) (int, error)
	CreateB2BPay(id int, input User_balance.B2BPay) (int, error)
}

type Balance interface {
	GetBalance(id int) (float32, error)
}

type History interface {
	GetHistory(id int) ([]User_balance.History, error)
}

type Repository struct {
	Authorization
	Payments
	Balance
	History
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Balance:       NewBalancePostgres(db),
		Payments:      NewPayPostgres(db),
		History:       NewHistoryPostgres(db),
	}
}
