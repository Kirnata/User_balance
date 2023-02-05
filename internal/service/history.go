package service

import (
	"github.com/Kirnata/User_balance"
	"github.com/Kirnata/User_balance/internal/repository"
)

type HistoryService struct {
	repo repository.History
}

func NewHistoryService(repo repository.History) *HistoryService {
	return &HistoryService{
		repo: repo,
	}
}

func (s *HistoryService) GetHistory(id int) ([]User_balance.History, error) {
	return s.repo.GetHistory(id)
}
