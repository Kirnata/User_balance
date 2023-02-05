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
	//CreateSinglePay(id int) (int, error)
}

type Balance interface {
	GetBalance(id int) (float32, error)
}

type Service struct {
	Authorization
	Payments
	Balance
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
		Balance:       NewBalanceService(repos.Balance),
	}
}
