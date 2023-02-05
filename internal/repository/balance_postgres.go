package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type BalanceServicePostgres struct {
	db *sqlx.DB
}

func NewBalancePostgres(db *sqlx.DB) *BalanceServicePostgres {
	return &BalanceServicePostgres{
		db: db,
	}
}

func (r *BalanceServicePostgres) GetBalance(id int) (float32, error) {
	var balance float32
	query := fmt.Sprintf("SELECT balance FROM %s WHERE id=$1", userTable)
	err := r.db.Get(&balance, query, id)
	return balance, err
}
