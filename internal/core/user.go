package core

import "errors"

type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInvitePostgres struct {
	Username string `json:"username" db:"username" binding:"required"`
}

func (u *UserInvitePostgres) Validate() error {
	if u.Username == "" {
		return errors.New("update payload has no required fields")
	}
	return nil
}
