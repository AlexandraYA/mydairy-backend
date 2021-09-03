package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	dairy "github.com/alexandraya/mydairy-backend"
	"github.com/alexandraya/mydairy-backend/pkg/repository"
)

const (
	salt       = "lsjfowierew75r3lmdfvi"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user dairy.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
