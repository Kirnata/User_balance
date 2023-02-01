package service

import (
	"github.com/Kirnata/User_balance"
	"github.com/Kirnata/User_balance/internal/repository"
)

type Authorization interface {
	CreateUser(user User_balance.User) (int, error)
}

type Transaction interface {
}

type Service struct {
	Authorization
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
	}
}
