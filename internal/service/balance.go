package service

import "github.com/Kirnata/User_balance/internal/repository"

type BalanceService struct {
	repo repository.Balance
}

func NewBalanceService(repo repository.Balance) *BalanceService {
	return &BalanceService{
		repo: repo,
	}
}

func (s *BalanceService) GetBalance(id int) (float32, error) {
	return s.repo.GetBalance(id)
}
