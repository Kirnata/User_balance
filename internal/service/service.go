package service

import (
	"github.com/Kirnata/User_balance"
	"github.com/Kirnata/User_balance/internal/repository"
)

type Authorization interface {
	CreateUser(user User_balance.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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

type Service struct {
	Authorization
	Payments
	Balance
	History
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
		Balance:       NewBalanceService(repos.Balance),
		Payments:      NewPayService(repos.Payments),
		History:       NewHistoryService(repos.History),
	}
}
