package service

import (
	"github.com/itoqsky/money-tracker-backend/internal/core"
	"github.com/itoqsky/money-tracker-backend/internal/storage"
)

type Authorization interface {
	CreateUser(user core.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
}

type Group interface {
	Create(userId int, group core.Group) (int, error)
	GetAll(userId int) ([]core.Group, error)
	GetById(userId, groupId int) (core.Group, error)
}

type Debt interface {
}

type Purchase interface {
}

type Service struct {
	Authorization
	User
	Group
	Debt
	Purchase
}

func NewService(store *storage.Storage) *Service {
	return &Service{
		Authorization: NewAuthService(store.Authorization),
		Group:         NewGroupService(store.Group),
	}
}
