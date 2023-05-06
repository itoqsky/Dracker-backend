package storage

import (
	"github.com/itoqsky/money-tracker-backend/internal/core"
	"github.com/jmoiron/sqlx"
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

type Storage struct {
	Authorization
	Debt
	Purchase
	User
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Authorization: NewAuthPostgres(db),
	}
}
