package repository

import (
	"fmt"
	"github.com/Kirnata/User_balance"
	"github.com/jmoiron/sqlx"
)

type PayPostgres struct {
	db *sqlx.DB
}

func NewPayPostgres(db *sqlx.DB) *PayPostgres {
	return &PayPostgres{
		db: db,
	}
}

func (r *PayPostgres) CreateSinglePay(id int, input User_balance.SinglePay) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	queryBalance := fmt.Sprintf("UPDATE %s SET balance = balance - $1  WHERE id=$2", userTable)
	if _, err := tx.Exec(queryBalance, input.Amount, id); err != nil {
		tx.Rollback()
		return 0, err
	}
	var noteId int
	queryHistory := fmt.Sprintf("INSERT INTO  %s (sender_id, value, description) VALUES ($1, $2, $3) RETURNING ID", historyTable)
	row := tx.QueryRow(queryHistory, id, input.Amount, input.Description)
	if err := row.Scan(&noteId); err != nil {
		tx.Rollback()
		return 0, err
	}
	return noteId, tx.Commit()
}

func (r *PayPostgres) CreateB2BPay(id int, input User_balance.B2BPay) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	queryToSend := fmt.Sprintf("UPDATE %s SET balance = balance - $1  WHERE id=$2", userTable)
	if _, err := tx.Exec(queryToSend, input.Amount, id); err != nil {
		tx.Rollback()
		return 0, err
	}
	queryToGet := fmt.Sprintf("UPDATE %s SET balance = balance + $1  WHERE id=$2", userTable)
	if _, err := tx.Exec(queryToGet, input.Amount, input.GetterId); err != nil {
		tx.Rollback()
		return 0, err
	}
	var noteId int
	queryHistory := fmt.Sprintf("INSERT INTO  %s (sender_id, getter_id, value, description) VALUES ($1, $2, $3, $4) RETURNING ID", historyTable)
	row := tx.QueryRow(queryHistory, id, input.GetterId, input.Amount, input.Description)
	if err := row.Scan(&noteId); err != nil {
		tx.Rollback()
		return 0, err
	}
	return noteId, tx.Commit()
}
