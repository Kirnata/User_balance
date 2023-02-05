package repository

import (
	"fmt"
	"github.com/Kirnata/User_balance"
	"github.com/jmoiron/sqlx"
)

type HistoryPostgres struct {
	db *sqlx.DB
}

func NewHistoryPostgres(db *sqlx.DB) *HistoryPostgres {
	return &HistoryPostgres{
		db: db,
	}
}

func (r *HistoryPostgres) GetHistory(id int) ([]User_balance.History, error) {
	var history []User_balance.History
	query := fmt.Sprintf("SELECT * FROM %s WHERE sender_id=$1 OR getter_id=$2", historyTable)
	err := r.db.Select(&history, query, id, id)
	if err != nil {
		return nil, err
	}
	return history, nil
}
