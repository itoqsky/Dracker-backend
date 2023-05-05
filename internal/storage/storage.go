package storage

import "github.com/jmoiron/sqlx"

type Authorization interface {
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

func NewStorage(*sqlx.DB) *Storage {
	return &Storage{}
}
