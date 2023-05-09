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
	GetAll(groupId int) ([]core.UserInputGetAll, error)
	Invite(groupId int, username string) error
	KickUser(gropId, kickUserId int) error
}

type Group interface {
	Create(userId int, group core.Group) (int, error)
	GetAll(userId int) ([]core.Group, error)
	GetById(userId, groupId int) (core.Group, error)
	Delete(users_sz, userId, gropId int) error
	Update(userId, groupId int, input core.UpdateGroupInput) error
}

type Debt interface {
}

type Purchase interface {
	Create(purchase core.Purchase, users []core.UserInputGetAll) (core.CreatePurchaseResponse, error)
	GetAll(groupId int) ([]core.Purchase, error)
	GetById(id int) (core.Purchase, error)
	Update(purchase core.Purchase, users []core.UserInputGetAll) error
	Delete(purchase core.Purchase, users []core.UserInputGetAll) error
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
		User:          NewUserPostgres(db),
		Purchase:      NewPurchasePostgres(db),
	}
}
