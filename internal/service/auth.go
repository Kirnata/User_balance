package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/Kirnata/User_balance"
	"github.com/Kirnata/User_balance/internal/repository"
	"time"
)

const (
	salt       = "sdfvjwe34foi3wegr5bjkwv5wervhi6wd"
	tokenTTL   = 12 * time.Hour
	signingKey = "ascjgsfdvqeqw;fovljvnkjnavwjw"
)

type AuthService struct {
	repo repository.Authorization
}

func newAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user User_balance.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
