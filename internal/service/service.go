package service

import (
	"github.com/itoqsky/money-tracker-backend/internal/core"
	"github.com/itoqsky/money-tracker-backend/internal/storage"
)

type Authorization interface {
	CreateUser(user core.User) (int, error)
}

type Debt interface {
}

type Purchase interface {
}

type User interface {
}

type Service struct {
	Authorization
	Debt
	Purchase
	User
}

func NewService(store *storage.Storage) *Service {
	return &Service{
		Authorization: NewAuthService(store.Authorization),
	}
}
