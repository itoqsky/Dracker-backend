package service

import "github.com/itoqsky/money-tracker-backend/internal/storage"

type Authorization interface {
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

func NewService(stor *storage.Storage) *Service {
	return &Service{}
}
