package service

import "github.com/Kirnata/User_balance/internal/repository"

type Authorization interface {
}

type Transaction interface {
}

type Service struct {
	Authorization
	Transaction
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
