package service

import (
	"github.com/Kirnata/User_balance"
	"github.com/Kirnata/User_balance/internal/repository"
)

type PayService struct {
	repo repository.Payments
}

func NewPayService(repo repository.Payments) *PayService {
	return &PayService{
		repo: repo,
	}
}

func (s PayService) CreateSinglePay(id int, input User_balance.SinglePay) (int, error) {
	return s.repo.CreateSinglePay(id, input)
}
