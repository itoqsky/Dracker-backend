package storage

import (
	"github.com/itoqsky/money-tracker-backend/internal/core"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user core.User) (int, error)
	GetUser(email, password string) (core.User, error)
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

type Storage struct {
	Authorization
	User
	Group
	Debt
	Purchase
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Authorization: NewAuthPostgres(db),
		Group:         NewGroupPostgres(db),
	}
}
