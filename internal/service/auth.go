package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/itoqsky/money-tracker-backend/internal/core"
	"github.com/itoqsky/money-tracker-backend/internal/storage"
)

const salt = "fe3fi9h208fnckdqnfie"

type AuthService struct {
	store storage.Authorization
}

func NewAuthService(store storage.Authorization) *AuthService {
	return &AuthService{store: store}
}

func (s *AuthService) CreateUser(user core.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.store.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
