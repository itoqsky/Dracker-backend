package core

import "errors"

type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInputGetAll struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

type UserInputKick struct {
	Id int `json:"id" binding:"required"`
}

type UserInputInvite struct {
	Username string `json:"username" binding:"required"`
}

func (u *UserInputInvite) Validate() error {
	if u.Username == "" {
		return errors.New("update payload has no required fields")
	}
	return nil
}
